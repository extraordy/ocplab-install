# OpenShift lab provisioning tool

### Overview
This repository contains a Terraform manifest file to deploy a set of nodes
for an OpenShift cluster and a Go binary tool to manage the installation in an
user-friendly way.

### Prerequisites
[**Terraform**](https://github.com/hashicorp/terraform) must be installed in the 
system. 
The [**Libvirt provider**](https://github.com/dmacvicar/terraform-provider-libvirt)
must be installed in the system.

The cluster is built upon customized RHEL7 images with enabled root access (safe
for test labs only) and cloud-init removed. Future releases will provide a 
custom cloud-init configuration.

The RHEL 7.6 KVM Guest Image can be downloaded 
[here](https://access.redhat.com/downloads/content/69/ver=/rhel---7/7.6/x86_64/product-software).
The latest CentOS KVM Guest image can be downloaded
[here](https://cloud.centos.org/centos/7/images/CentOS-7-x86_64-GenericCloud.qcow2)

After downloading the Guest Image a little workaround is needed to modify its 
content using the tool **virt-customize**, installed by the **libguestfs-tools**
package.  
The example below updates the root password to **redhat** and removes cloud-init.

```
$ virt-customize -a /path/to/image.qcow --root-password password:redhat --uninstall cloud-init
```

The image should be available in a path available to Libvirt, usually a volume
pool.

### Build
To build and install the setup tool:
```
$ make
$ sudo make install
```
The installation tool will be installed in **/usr/local/bin/ocplab-install**.

### Usage
Initialize the Terraform working directory:
```
$ ocplab-install init
```

Create the cluster nodes:
```
$ ocplab-install create
```

To provide a custom source image path:
```
$ ocplab-install create --img /foo/bar
```

To destroy all the resources:
```
$ ocplab-install destroy
```

### TODO
- Provide a pre-customized image
- Improve image management (maybe add pull from the tool)
- Configure predefined IPs for instances
- Configure outputs to print out post-install informations
- Configure ssh keys to inject into the instances (using cloud-init)
- Use the **github.com/hashicorp/terraform/pkg/command** and **github.com/hashicorp/go-plugin** 
  packages instead of the external binary.
- Use the **github.com/dmacvicar/terraform-provider-libvirt/libvirt** package instead of
  the external binary.

### Maintainer
- Giovan Battista Salinetti <gbsalinetti@extraordy.com>

