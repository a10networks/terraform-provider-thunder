---
layout: "thunder"
page_title: "thunder: thunder_virtual_server"
sidebar_current: "docs-thunder-resource-virtual-server"
description: |-
    Provides details about thunder virtual server resource for A10
---

# thunder\_virtual\_server

`thunder_virtual_server` provides details about configuring virtual server on device
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_virtual_server" "vs9" {
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
