packer {
  required_plugins {
    libvirt = {
      version = ">= 1.0.0"
      source  = "local/homemade/libvirt"
    }
  }
}

variable "root_password" {
  type    = string
  description = "The root password"
}

source "libvirt" "xcpng" {
  libvirt_uri = "qemu:///system"

  vcpu   = 2
  memory = 4096
  
  network_interface {
    type  = "managed"
    alias = "communicator"
  }
  network_address_source = "lease"

  communicator {
    communicator         = "ssh"
    ssh_username         = "root"
    ssh_password         = "DefaultRootPasswordUsedInAnswerfile"
    ssh_timeout          = "10m"
  }
  
  graphics {
    type = "vnc"
  }
  boot_devices = ["hd", "cdrom"]

  volume {
    device = "cdrom"
    pool   = "default"
    name   = "xcpng.iso"
    source {
      type     = "external"
      urls     = ["https://xcpng-mirror.as208069.net/isos/8.3/xcp-ng-8.3.0.iso"]
      checksum = "49b6143d1bbb1fd0bb1f7ce873817190c1f2ca394c258ce7b1f8e13a75c72674"
    }
    bus    = "sata"
    format = "raw"
  }

  volume {
    alias = "artifact"

    pool  = "default"
    name  = "xcpng.qcow2"

    format   = "qcow2"
    capacity = "100G"
    bus      = "sata"
  }

  http_directory = "files/http"
  boot_command = [
    "<esc><wait5>",
    "<f2><wait2>",
    "menu<enter><wait2>",
    "<tab><wait2>",
    "<end><left><left><left><left><left><left><left><left><left><left><left><left><left><left><left><left><wait2>",
    "answerfile=http://{{ .HTTPIP }}:{{ .HTTPPort }}/answerfile install <enter>",
  ]
}

build {
  sources = ["source.libvirt.xcpng"]

  provisioner "file" {
    source      = "files/minimal_connectivity.bash"
    destination = "/usr/local/bin/minimal_connectivity.bash"
  }
  provisioner "file" {
    source      = "files/minimal_connectivity.service"
    destination = "/etc/systemd/system/minimal_connectivity.service"
  }
  provisioner "file" {
    source      = "files/sshd_config"
    destination = "/etc/ssh/sshd_config"
  }

  provisioner "shell" {
    inline = [
      "echo '${var.root_password}' | passwd --stdin root", #BEWARE IT WILL BREAK FURTHER SSH CONNECTION
      "echo \"add_drivers+=' nvme '\" > /etc/dracut.conf.d/nvme.conf",
      "dracut --force -f /boot/initrd-4.19.0+1.img",
      "chmod 600 /etc/ssh/sshd_config",
      "chmod 755 /usr/local/bin/minimal_connectivity.bash",
      "systemctl enable minimal_connectivity.service",
      "systemctl daemon-reload",
      "nohup bash -c 'sleep 10; poweroff' &",
      "rm -f /tmp/setup.bash",
    ]
    remote_path = "/tmp/setup.bash"
    skip_clean = true # As ssh connections are broken
  }
}
