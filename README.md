Terraform Provider
=========================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Maintainers
-----------

This provider plugin is maintained by the Terraform team at [HashiCorp](https://www.hashicorp.com/).

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.12 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-vthunder`

```sh
$ git clone git@github.com:terraform-providers/terraform-provider-vthunder $GOPATH/src/github.com/terraform-providers/terraform-provider-vthunder
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-vthunder
$ make build
```

# Usage

### Provider Configuration

```
provider "vthunder" {
address = "129.213.86.193"
username = "myUser"  
password = "myPassword“
}
```
##### Argument Reference
The following arguments are supported.
•	username - User name to access vThunder.
•	password - Password to access vThunder.
•	address – IP address of vThunder; to be configured.


### Resource Configuration

#### vthunder_ethernet

```
resource "vthunder_ethernet" "eth"{
ethernet_list{
ifnum=1
ip{
address_list={
ipv4_address="10.0.2.9"
ipv4_netmask="255.255.255.0"
}
}
action="enable"
}
ethernet_list{
ifnum=2
ip{
#dhcp=1
address_list={
ipv4_address="10.0.3.9"
ipv4_netmask="255.255.255.0"
}
}
action="enable"
}
}
```

#### Argument Reference
See https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-GR1-P1/html/axapiv3/interface.html for possible values for these arguments and for an exhaustive list of arguments.

#### vthuder_rib_route
```
resource "vthunder_rib_route" "rib"{
ip_dest_addr="0.0.0.0"
ip_mask="/0"
ip_nexthop_ipv4{
ip_next_hop="10.0.2.9"
distance_nexthop_ip=1
}
}
```

#### Argument Reference
See https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-GR1-P1/html/axapiv3/ip_route_rib.html for possible values for these arguments and for an exhaustive list of arguments.

### vthunder_server
```
resource "vthunder_server" "rs9" {
health_check_disable=1
name="rs9"
host="10.0.3.8"
port_list {
health_check_disable=1
port_number=80
protocol="tcp"
}
}
```

#### Argument Reference
See https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-GR1-P1/html/axapiv3/slb_server.html for possible values for these arguments and for an exhaustive list of arguments.

### vthunder_service_group
```
resource "vthunder_service_group" "sg9" {
name="sg9"
protocol="TCP"
member_list {
name="${vthunder_server.rs9.name}"
port=80
}
}
```

#### Argument Reference
See https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-GR1-P1/html/axapiv3/slb_service_group.html for possible values for these arguments and for an exhaustive list of arguments. Backend server name needs to be specified in the member list. This name can come from vthunder_server resource at runtime as specified in the sample.

### vthunder_virtual_server
```
resource "vthunder_virtual_server" "vs9" {
name="vs9"
ha_dynamic = 1
ip_address="10.0.2.7"
port_list {
auto=1
port_number=8080
protocol="tcp"
service_group="${vthunder_service_group.sg9.name}"
snat_on_vip=1
}
}
```

#### Argument Reference
See https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-GR1-P1/html/axapiv3/slb_virtual_server.html for possible values for these arguments and for an exhaustive list of arguments.

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make bin
...
$ $GOPATH/bin/terraform-provider-vthunder
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

## Samples
See the examples directory for various LB topologies that can be driven from this terraform provider.

## Bug Reporting and Feature Requests
Please submit bug reports and feature requests via GitHub issues. When reporting bugs, please include the Terraform script that demonstrates the bug and the command output. Stack traces will be helpful. Please ensure any sensitive information is redacted as Issues and Pull Requests are publicly viewable.

## Contact
If you have a question that cannot be submitted via Github Issues, please email support@a10networks.com with "a10-terraform-provider" in the subject line.
