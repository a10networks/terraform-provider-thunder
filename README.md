# terraform-provider-vThunder
Terraform Provider for A10 vThunder

# Description
This project is a Terraform custom provider for A10's vThunder device. It uses the aXAPIs to create/configure LB configurations.

# Requirement

- [hashicorp/terraform](https://github.com/hashicorp/terraform/)

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

## Building
### Assumption
•	You have some experience with Terraform, the different provisioners and providers that come out of the box, its configuration files, tfstate files, etc.
•	You are comfortable with the Go language and its code organization.
1.	Install terraform from https://www.terraform.io/downloads.html
2.	Install dep (https://github.com/golang/dep)
3.	Check out this code: git clone https://<>
4.	Build this code using
    ```        
    go build -o terraform-provider-vthunder
    ```
5.  Copy the binary (either from the build or from the releases page) terraform-provider-vthunder to an appropriate location.
2.	Run Terraform as usual
    ```
    terraform init -plugin-dir="<directory path of terraform-provider-vthunder binary>"
    terraform plan
    terraform apply
    ```
## Samples
See the examples directory for various LB topologies that can be driven from this terraform provider.

## Bug Reporting and Feature Requests
Please submit bug reports and feature requests via GitHub issues. When reporting bugs, please include the Terraform script that demonstrates the bug and the command output. Stack traces will be helpful. Please ensure any sensitive information is redacted as Issues and Pull Requests are publicly viewable.

## Contact
If you have a question that cannot be submitted via Github Issues, please email support@a10networks.com with "a10-terraform-provider" in the subject line.

