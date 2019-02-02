# OpenShift lab provisioning tool

### Prerequisites
[**Terraform**](https://github.com/hashicorp/terraform) must be installed in the 
system. 
The [**Libvirt provider**](https://github.com/dmacvicar/terraform-provider-libvirt)
must be installed in the system.

The cluster is built upon customized RHEL7 images with enabled root access (safe
for test labs only) and cloud-init removed. Future releases will provide a 
custom cloud-init configuration.
After downloading the RHEL 7.6 KVM Guest Image modify its content using the
tool **virt-customize**:

```
$ virt-customize -a /path/to/image.qcow --root-password password:redhat --uninstall cloud-init
```

The example above updates the root password to **redhat** and removes cloud-init.

### Build
To build and install the setup tool:
```
$ make
$ sudo make install
```

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
- Improve image management
- Configure predefined IPs
- Configure outputs
- Configure ssh keys

### Maintainer
- Giovan Battista Salinetti <gbsalinetti@extraordy.com>

