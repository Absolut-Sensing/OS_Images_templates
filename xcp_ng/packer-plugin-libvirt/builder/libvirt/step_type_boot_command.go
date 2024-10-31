package libvirt

import (
	"context"
	"log"
	"strings"

	"github.com/digitalocean/go-libvirt"
	"github.com/hashicorp/packer-plugin-sdk/bootcommand"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
	"golang.org/x/mobile/event/key"

	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

type templateArgs struct {
	HTTPIP   string
	HTTPPort int
}

func templatizeCommands(strings []string, ctx *interpolate.Context) ([]string, error) {
	var err error

	if strings == nil {
		return make([]string, 0), err
	}

	templatedStrings := make([]string, len(strings))
	for index, string := range strings {
		templatedStrings[index], err = interpolate.Render(string, ctx)
		if err != nil {
			return nil, err
		}
	}

	return templatedStrings, err
}

type stepTypeBootCommand struct{}

func (s *stepTypeBootCommand) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)
	driver := state.Get("driver").(*libvirt.Libvirt)
	domain := state.Get("domain").(*libvirt.Domain)
	config := state.Get("config").(*Config)

	if len(config.BootConfig.BootCommand) == 0 {
		log.Printf("Skip typing boot command as boot command is empty")
	}

	httpIp := state.Get("http_ip").(string)
	httpPort := state.Get("http_port").(int)
	ictx := config.ctx
	ictx.Data = templateArgs{
		HTTPIP:   httpIp,
		HTTPPort: httpPort,
	}

	var sendKeys bootcommand.SendUsbScanCodes = func(k key.Code, down bool) error {
		var keys []uint32
		if down {
			keys = append(keys, uint32(key.CodeLeftShift)) // Append the Shift key if needed (must be BEFORE key)
		}
		keys = append(keys, uint32(k))

		err := driver.DomainSendKey(
			*domain,
			uint32(libvirt.KeycodeSetUsb),
			uint32(config.BootConfig.KeyHoldType.Milliseconds()),
			keys,
			0,
		)

		if err != nil {
			return err
		}
		return nil
	}

	// Interpolate each string in qemuargs
	templatedBootCommand, err := templatizeCommands(config.BootConfig.BootCommand, &ictx)
	if err != nil {
		return haltOnError(ui, state, "%s", err)
	}
	flatBootCommand := strings.Join(templatedBootCommand, "")
	seq, err := bootcommand.GenerateExpressionSequence(flatBootCommand)

	if err != nil {
		return haltOnError(ui, state, "%s", err)
	}

	err = seq.Do(ctx, bootcommand.NewUSBDriver(sendKeys, 0))

	if err != nil {
		return haltOnError(ui, state, "%s", err)
	}

	return multistep.ActionContinue
}

func (s *stepTypeBootCommand) Cleanup(bag multistep.StateBag) {
	// Do nothing
}
