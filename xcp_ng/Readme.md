# Install libvirt :
  ...
  echo '"/var/lib/libvirt/images/" r,' | sudo tee -a /etc/apparmor.d/local/abstractions/libvirt-qemu >/dev/null
  echo '"/var/lib/libvirt/images/**" rwk,' | sudo tee -a /etc/apparmor.d/local/abstractions/libvirt-qemu >/dev/null

# Install packer plugin
This is a slightly modified packer version with support of http server & allowing shift usage in boot command.
https://developer.hashicorp.com/packer/tutorials/docker-get-started/get-started-install-cli

```bash
(
  cd xcp_ng/packer-plugin-libvirt/
  go build
  plugin_folder="$HOME/.config/packer/plugins/local/homemade/libvirt"
  plugin_file="${plugin_folder}/packer-plugin-libvirt_v1.0.0_x5.0_linux_amd64"
  plugin_sha256="${plugin_folder}/packer-plugin-libvirt_v1.0.0_x5.0_linux_amd64_SHA256SUM"
  mkdir "${plugin_folder}"
  cp packer-plugin-libvirt "${plugin_file}"
  sha256sum "${plugin_file}" | awk '{print $1}' > "${plugin_sha256}"
)
```

```bash
(
  cd xcp-ng
  packer init .
  packer build -var 'root_password=<value>' .
)
```