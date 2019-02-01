# OpenShift lab provisioning tool

### Prerequisites
[**Terraform**](https://github.com/hashicorp/terraform) must be installed in the 
system. 
The [**Libvirt provider**](https://github.com/dmacvicar/terraform-provider-libvirt)
must be installed in the system.

### Build
To build and install the setup tool:
```
$ make
$ sudo make install
```

### Usage
Initialize the Terraform working directory:
```
$ ocplab-setup init
```

Create the cluster nodes:
```
$ ocplab-setup create
```

To destroy all the resources:
```
$ ocplab-setup destroy
```

### TODO
- Configure predefined IPs
- Configure outputs
- Configure ssh keys

### Maintainer
- Giovan Battista Salinetti <gbsalinetti@extraordy.com>

