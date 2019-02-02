################################################################################
#
# Provider section
#
################################################################################

provider "libvirt" {
    uri = "qemu:///system"
}

################################################################################
#
# Vars section
#
################################################################################

variable "master-count" {
    default = 3
}

variable "infra-count" {
    default = 3
}

variable "worker-count" {
    default = 3
}

variable "lb-count" {
    default = 1
}

variable "source-image" {
    default = "/home/qemu/libvirt/default/rhel-server-7.6-x86_64-kvm.qcow2"
}

variable "disk0-size" {
    default = 10737418240
}

################################################################################
#
# Network section
#
################################################################################

resource "libvirt_network" "cluster_net" {
    name = "cluster_net"
    addresses = ["10.10.1.0/24"]
    mode = "nat"
    domain = "extlab.io"
    dhcp {
        enabled = true
    }
}

################################################################################
#
# Volumes section
#
################################################################################

resource "libvirt_volume" "os_image" {
    name = "os_image"
    pool = "default"
    source = "${var.source-image}"
}

resource "libvirt_volume" "master-disk0" {
    count = "${var.master-count}"
    name = "${format("ocplabm-%02d-disk0", count.index + 1)}"
    pool = "default"
    format = "qcow2"
    size = "${var.disk0-size}"
    base_volume_id = "${libvirt_volume.os_image.id}"
}

resource "libvirt_volume" "infra-disk0" {
    count = "${var.infra-count}"
    name = "${format("ocplabi-%02d-disk0", count.index + 1)}"
    pool = "default"
    format = "qcow2"
    size = "${var.disk0-size}"
    base_volume_id = "${libvirt_volume.os_image.id}"
}


resource "libvirt_volume" "worker-disk0" {
    count = "${var.worker-count}"
    name = "${format("ocplabw-%02d-disk0", count.index + 1)}"
    pool = "default"
    format = "qcow2"
    size = "${var.disk0-size}"
    base_volume_id = "${libvirt_volume.os_image.id}"
}

resource "libvirt_volume" "lb-disk0" {
    count = "${var.lb-count}"
    name = "${format("ocplablb-%02d-disk0", count.index + 1)}"
    pool = "default"
    format = "qcow2"
    size = "${var.disk0-size}"
    base_volume_id = "${libvirt_volume.os_image.id}"
}

################################################################################
#
# Instances section
#
################################################################################

resource "libvirt_domain" "ocpmaster" {
    count = "${var.master-count}"
    name = "${format("ocplabm-%02d", count.index + 1)}"
    memory = 4096
    vcpu = 1
    disk {
        volume_id = "${element(libvirt_volume.master-disk0.*.id, count.index)}"
    }
    network_interface {
        network_id = "${libvirt_network.cluster_net.id}"
        hostname = "${format("ocplabm-%02d.extlab.io", count.index + 1)}"
        wait_for_lease = true
    }
    console {
        type = "pty"
        target_port = "0"
        target_type = "serial"
    }
    console {
        type = "pty"
        target_port = "1"
        target_type = "virtio"
    }
    graphics {
        type        = "spice"
        listen_type = "address"
        autoport    = true
    }
}

resource "libvirt_domain" "ocpinfra" {
    count = "${var.infra-count}"
    name = "${format("ocplabi-%02d", count.index + 1)}"
    memory = 4096
    vcpu = 1
    disk {
        volume_id = "${element(libvirt_volume.infra-disk0.*.id, count.index)}"
    }
    network_interface {
        network_id = "${libvirt_network.cluster_net.id}"
        hostname = "${format("ocplabi-%02d.extlab.io", count.index + 1)}"
        wait_for_lease = true
    }
    console {
        type = "pty"
        target_port = "0"
        target_type = "serial"
    }
    console {
        type = "pty"
        target_port = "1"
        target_type = "virtio"
    }
    graphics {
        type        = "spice"
        listen_type = "address"
        autoport    = true
    }
}

resource "libvirt_domain" "ocpworker" {
    count = "${var.worker-count}"
    name = "${format("ocplabw-%02d", count.index + 1)}"
    memory = 4096
    vcpu = 1
    disk {
        volume_id = "${element(libvirt_volume.worker-disk0.*.id, count.index)}"
    }
    network_interface {
        network_id = "${libvirt_network.cluster_net.id}"
        hostname = "${format("ocplabw-%02d.extlab.io", count.index + 1)}"
        wait_for_lease = true
    }
    console {
        type = "pty"
        target_port = "0"
        target_type = "serial"
    }
    console {
        type = "pty"
        target_port = "1"
        target_type = "virtio"
    }
    graphics {
        type        = "spice"
        listen_type = "address"
        autoport    = true
    }
}

resource "libvirt_domain" "ocplb" {
    count = "${var.lb-count}"
    name = "${format("ocplablb-%02d", count.index + 1)}"
    memory = 1024
    vcpu = 1
    disk {
        volume_id = "${element(libvirt_volume.lb-disk0.*.id, count.index)}"
    }
    network_interface {
        network_id = "${libvirt_network.cluster_net.id}"
        hostname = "${format("ocplablb-%02d.extlab.io", count.index + 1)}"
        wait_for_lease = true
    }
    console {
        type = "pty"
        target_port = "0"
        target_type = "serial"
    }
    console {
        type = "pty"
        target_port = "1"
        target_type = "virtio"
    }
    graphics {
        type        = "spice"
        listen_type = "address"
        autoport    = true
    }
}
################################################################################
#
# Output section
#
################################################################################

