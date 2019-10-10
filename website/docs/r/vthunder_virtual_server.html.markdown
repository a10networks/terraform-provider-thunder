---
layout: "vthunder"
page_title: "vthunder: vthunder_virtual_server"
sidebar_current: "docs-vthunder-resource-virtual-server"
description: |-
    Provides details about vthunder virtual server resource for A10
---

# vthunder\_virtual\_server

`vthunder_virtual_server` provides details about configuring virtual server on device
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_virtual_server" "vs9" {
  name="vs9"
  ip_address="10.0.2.10"
  port_list {
    auto=1
    port_number=8080
    protocol="tcp"
    service_group="sg9"
    snat_on_vip=1
  }
}
```

## Argument Reference

* `name` - Name of virtual server
* `ip_address` - IP address
* `auto` - Configure auto NAT for the vport
* `port_number` - Port number
* `protocol` - protocol
* `service_group` - Service group name
* `snat_on_vip` - Enable source NAT traffic against VIP
