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

resource "thunder_ip_nat_pool" "test_pool" {
        
        pool_name = "SNAT_VRID1"
        start_address = "1.1.1.1"
        end_address = "1.1.1.10"
        netmask = "/24"
        vrid = 4
        ip_rr = 1
        port_overload = 1        
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

