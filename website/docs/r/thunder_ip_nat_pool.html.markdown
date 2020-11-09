---
layout: "thunder"
page_title: "thunder: thunder_ip_nat_pool"
sidebar_current: "docs-thunder-resource-ip-nat-pool"
description: |-
	Provides details about thunder ip nat pool resource for A10
---

# thunder\_ip\_nat\_pool

`thunder_ip_nat_pool` Provides details about thunder ip nat pool
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_ip_nat_pool" "Ip_Nat_Pool_Test" {

use_if_ip = 0
uuid = "string"
start_address = "string"
port_overload = 0
vrid = 0
netmask = "string"
end_address = "string"
ip_rr = 0
chunk_netmask = "string"
ethernet = 0
scaleout_device_id = 0
gateway = "string"
pool_name = "string"
 
}

```

## Argument Reference

* `end_address` - Configure end IP address of NAT pool
* `ethernet` - Ethernet interface
* `gateway` - Configure gateway IP
* `ip_rr` - Use IP address round-robin behavior
* `netmask` - Configure mask for pool
* `pool_name` - Specify pool name or pool group
* `scaleout_device_id` - Configure Scaleout device id to which this NAT pool is to be bound (Specify Scaleout device id)
* `start_address` - Configure start IP address of NAT pool
* `use_if_ip` - Use Interface IP
* `uuid` - uuid of the object
* `vrid` - Configure VRRP-A vrid (Specify ha VRRP-A vrid)
