#!/bin/bash
set -o errexit   # abort on nonzero exit status
set -o nounset   # abort on unbound variable
set -o pipefail  # don't hide errors within pipes, Bash only

exec > >(tee -a /var/log/minimal_connectivity) 2>&1

echo "=== Minimal Connectivity Start ==="

# Read user_data
echo "Determining root filesystem and metadata partition..."
rootfs_partiton_path="$(mount | grep -E '^/dev/[a-zA-Z0-9]+ on / ' | awk '{print $1}')"
rootfs_disk="$(echo "${rootfs_partiton_path}" | sed -E 's|^/dev/(.+)p[0-9]+$|\1|')"
metadata_partition="$(find /dev/ -name "${rootfs_disk}p*" | sort | tail -n 1)"
mount "${metadata_partition}" /mnt/
echo "Mounted metadata partition ${metadata_partition} on mountpoint /mnt/."

echo -e "Userdata : \n$(cat /mnt/openstack/latest/user_data)"
source /mnt/openstack/latest/user_data # WARNING: Sourcing external data can be risky

# Set XenServer Management Interface
echo "Configuring XenServer management interface..."
while ! xe host-list --minimal; do # Seems that waiting for xapi.target if not enough for xen to be up & ready
    echo not yet ready;
    sleep 5; 
done;
host_uuid="$(xe host-list name-label="$(hostname)" params=uuid --minimal)"
xe pif-scan host-uuid="${host_uuid}"
old_mgmt_pif_uuid="$(xe pif-list management=true params=uuid --minimal)"
new_mgmt_pif_uuid="$(xe pif-list MAC="${MGMT_IFACE_MAC}" params=uuid --minimal)"

if [[ "${new_mgmt_pif_uuid}" != "${old_mgmt_pif_uuid}" ]]; then
    echo "Changing management interface from ${old_mgmt_pif_uuid:-None} to ${new_mgmt_pif_uuid}..."
    xe host-management-disable
    xe pif-reconfigure-ip uuid="${new_mgmt_pif_uuid}" mode=dhcp
    xe host-management-reconfigure pif-uuid="${new_mgmt_pif_uuid}"
    [ -n "${old_mgmt_pif_uuid}" ] && xe pif-forget uuid="${old_mgmt_pif_uuid}"
    xe pif-scan host-uuid="${host_uuid}"
fi

# Add Authorized SSH Key
if ! grep -q "${AUTHORIZED_SSH_KEY}" /root/.ssh/authorized_keys; then
    echo "Adding authorized SSH key..."
    mkdir -p /root/.ssh/ && touch /root/.ssh/authorized_keys
    echo "${AUTHORIZED_SSH_KEY}" >> /root/.ssh/authorized_keys
fi

# Unmount Metadata Partition
umount "${metadata_partition}"
sudo systemctl disable minimal_connectivity.service

echo "=== Minimal Connectivity End ==="
