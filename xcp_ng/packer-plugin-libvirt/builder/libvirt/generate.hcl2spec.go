// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package libvirt

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/communicator"
	"github.com/thomasklein94/packer-plugin-libvirt/builder/libvirt/network"
	"github.com/thomasklein94/packer-plugin-libvirt/builder/libvirt/volume"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName       *string                        `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType     *string                        `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion     *string                        `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug           *bool                          `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce           *bool                          `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError         *string                        `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars        map[string]string              `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars   []string                       `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	HTTPDir               *string                        `mapstructure:"http_directory" cty:"http_directory" hcl:"http_directory"`
	HTTPContent           map[string]string              `mapstructure:"http_content" cty:"http_content" hcl:"http_content"`
	HTTPPortMin           *int                           `mapstructure:"http_port_min" cty:"http_port_min" hcl:"http_port_min"`
	HTTPPortMax           *int                           `mapstructure:"http_port_max" cty:"http_port_max" hcl:"http_port_max"`
	HTTPAddress           *string                        `mapstructure:"http_bind_address" cty:"http_bind_address" hcl:"http_bind_address"`
	HTTPInterface         *string                        `mapstructure:"http_interface" undocumented:"true" cty:"http_interface" hcl:"http_interface"`
	Communicator          *communicator.FlatConfig       `mapstructure:"communicator" cty:"communicator" hcl:"communicator"`
	BootGroupInterval     *string                        `mapstructure:"boot_keygroup_interval" cty:"boot_keygroup_interval" hcl:"boot_keygroup_interval"`
	BootWait              *string                        `mapstructure:"boot_wait" cty:"boot_wait" hcl:"boot_wait"`
	BootCommand           []string                       `mapstructure:"boot_command" cty:"boot_command" hcl:"boot_command"`
	KeyHoldType           *string                        `mapstructure:"key_hold_time" cty:"key_hold_time" hcl:"key_hold_time"`
	DomainName            *string                        `mapstructure:"domain_name" required:"false" cty:"domain_name" hcl:"domain_name"`
	MemorySize            *int                           `mapstructure:"memory" required:"false" cty:"memory" hcl:"memory"`
	CpuCount              *int                           `mapstructure:"vcpu" required:"false" cty:"vcpu" hcl:"vcpu"`
	CpuMode               *string                        `mapstructure:"cpu_mode" required:"false" cty:"cpu_mode" hcl:"cpu_mode"`
	NetworkInterfaces     []network.FlatNetworkInterface `mapstructure:"network_interface" required:"false" cty:"network_interface" hcl:"network_interface"`
	CommunicatorInterface *string                        `mapstructure:"communicator_interface" required:"false" cty:"communicator_interface" hcl:"communicator_interface"`
	Volumes               []volume.FlatVolume            `mapstructure:"volume" required:"false" cty:"volume" hcl:"volume"`
	ArtifactVolumeAlias   *string                        `mapstructure:"artifact_volume_alias" required:"false" cty:"artifact_volume_alias" hcl:"artifact_volume_alias"`
	BootDevices           []string                       `mapstructure:"boot_devices" required:"false" cty:"boot_devices" hcl:"boot_devices"`
	DomainGraphics        []FlatDomainGraphic            `mapstructure:"graphics" required:"false" cty:"graphics" hcl:"graphics"`
	NetworkAddressSource  *string                        `mapstructure:"network_address_source" required:"false" cty:"network_address_source" hcl:"network_address_source"`
	LibvirtURI            *string                        `mapstructure:"libvirt_uri" required:"true" cty:"libvirt_uri" hcl:"libvirt_uri"`
	ShutdownMode          *string                        `mapstructure:"shutdown_mode" required:"false" cty:"shutdown_mode" hcl:"shutdown_mode"`
	ShutdownTimeout       *string                        `mapstructure:"shutdown_timeout" required:"false" cty:"shutdown_timeout" hcl:"shutdown_timeout"`
	DomainType            *string                        `mapstructure:"domain_type" required:"false" cty:"domain_type" hcl:"domain_type"`
	Arch                  *string                        `mapstructure:"arch" required:"false" cty:"arch" hcl:"arch"`
	Chipset               *string                        `mapstructure:"chipset" required:"false" cty:"chipset" hcl:"chipset"`
	LoaderPath            *string                        `mapstructure:"loader_path" required:"false" cty:"loader_path" hcl:"loader_path"`
	LoaderType            *string                        `mapstructure:"loader_type" required:"false" cty:"loader_type" hcl:"loader_type"`
	SecureBoot            *bool                          `mapstructure:"secure_boot" required:"false" cty:"secure_boot" hcl:"secure_boot"`
	NvramPath             *string                        `mapstructure:"nvram_path" required:"false" cty:"nvram_path" hcl:"nvram_path"`
	NvramTemplate         *string                        `mapstructure:"nvram_template" required:"false" cty:"nvram_template" hcl:"nvram_template"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":        &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"http_directory":             &hcldec.AttrSpec{Name: "http_directory", Type: cty.String, Required: false},
		"http_content":               &hcldec.AttrSpec{Name: "http_content", Type: cty.Map(cty.String), Required: false},
		"http_port_min":              &hcldec.AttrSpec{Name: "http_port_min", Type: cty.Number, Required: false},
		"http_port_max":              &hcldec.AttrSpec{Name: "http_port_max", Type: cty.Number, Required: false},
		"http_bind_address":          &hcldec.AttrSpec{Name: "http_bind_address", Type: cty.String, Required: false},
		"http_interface":             &hcldec.AttrSpec{Name: "http_interface", Type: cty.String, Required: false},
		"communicator":               &hcldec.BlockSpec{TypeName: "communicator", Nested: hcldec.ObjectSpec((*communicator.FlatConfig)(nil).HCL2Spec())},
		"boot_keygroup_interval":     &hcldec.AttrSpec{Name: "boot_keygroup_interval", Type: cty.String, Required: false},
		"boot_wait":                  &hcldec.AttrSpec{Name: "boot_wait", Type: cty.String, Required: false},
		"boot_command":               &hcldec.AttrSpec{Name: "boot_command", Type: cty.List(cty.String), Required: false},
		"key_hold_time":              &hcldec.AttrSpec{Name: "key_hold_time", Type: cty.String, Required: false},
		"domain_name":                &hcldec.AttrSpec{Name: "domain_name", Type: cty.String, Required: false},
		"memory":                     &hcldec.AttrSpec{Name: "memory", Type: cty.Number, Required: false},
		"vcpu":                       &hcldec.AttrSpec{Name: "vcpu", Type: cty.Number, Required: false},
		"cpu_mode":                   &hcldec.AttrSpec{Name: "cpu_mode", Type: cty.String, Required: false},
		"network_interface":          &hcldec.BlockListSpec{TypeName: "network_interface", Nested: hcldec.ObjectSpec((*network.FlatNetworkInterface)(nil).HCL2Spec())},
		"communicator_interface":     &hcldec.AttrSpec{Name: "communicator_interface", Type: cty.String, Required: false},
		"volume":                     &hcldec.BlockListSpec{TypeName: "volume", Nested: hcldec.ObjectSpec((*volume.FlatVolume)(nil).HCL2Spec())},
		"artifact_volume_alias":      &hcldec.AttrSpec{Name: "artifact_volume_alias", Type: cty.String, Required: false},
		"boot_devices":               &hcldec.AttrSpec{Name: "boot_devices", Type: cty.List(cty.String), Required: false},
		"graphics":                   &hcldec.BlockListSpec{TypeName: "graphics", Nested: hcldec.ObjectSpec((*FlatDomainGraphic)(nil).HCL2Spec())},
		"network_address_source":     &hcldec.AttrSpec{Name: "network_address_source", Type: cty.String, Required: false},
		"libvirt_uri":                &hcldec.AttrSpec{Name: "libvirt_uri", Type: cty.String, Required: false},
		"shutdown_mode":              &hcldec.AttrSpec{Name: "shutdown_mode", Type: cty.String, Required: false},
		"shutdown_timeout":           &hcldec.AttrSpec{Name: "shutdown_timeout", Type: cty.String, Required: false},
		"domain_type":                &hcldec.AttrSpec{Name: "domain_type", Type: cty.String, Required: false},
		"arch":                       &hcldec.AttrSpec{Name: "arch", Type: cty.String, Required: false},
		"chipset":                    &hcldec.AttrSpec{Name: "chipset", Type: cty.String, Required: false},
		"loader_path":                &hcldec.AttrSpec{Name: "loader_path", Type: cty.String, Required: false},
		"loader_type":                &hcldec.AttrSpec{Name: "loader_type", Type: cty.String, Required: false},
		"secure_boot":                &hcldec.AttrSpec{Name: "secure_boot", Type: cty.Bool, Required: false},
		"nvram_path":                 &hcldec.AttrSpec{Name: "nvram_path", Type: cty.String, Required: false},
		"nvram_template":             &hcldec.AttrSpec{Name: "nvram_template", Type: cty.String, Required: false},
	}
	return s
}

// FlatDomainGraphic is an auto-generated flat version of DomainGraphic.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatDomainGraphic struct {
	Type    *string `mapstructure:"type" required:"true" cty:"type" hcl:"type"`
	Display *string `mapstructure:"display" required:"false" cty:"display" hcl:"display"`
	Port    *int    `mapstructure:"port" required:"false" cty:"port" hcl:"port"`
}

// FlatMapstructure returns a new FlatDomainGraphic.
// FlatDomainGraphic is an auto-generated flat version of DomainGraphic.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*DomainGraphic) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatDomainGraphic)
}

// HCL2Spec returns the hcl spec of a DomainGraphic.
// This spec is used by HCL to read the fields of DomainGraphic.
// The decoded values from this spec will then be applied to a FlatDomainGraphic.
func (*FlatDomainGraphic) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"type":    &hcldec.AttrSpec{Name: "type", Type: cty.String, Required: false},
		"display": &hcldec.AttrSpec{Name: "display", Type: cty.String, Required: false},
		"port":    &hcldec.AttrSpec{Name: "port", Type: cty.Number, Required: false},
	}
	return s
}
