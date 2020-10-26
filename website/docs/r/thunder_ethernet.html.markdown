---
layout: "thunder"
page_title: "thunder: thunder_ethernet"
sidebar_current: "docs-thunder-resource-ethernet"
description: |-
    Provides details about thunder ethernet resource for A10
---

# thunder\_ethernet

`thunder_ethernet` provides details about how to enable ethernet interfaces on a device
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_ethernet" "eth"{
  ethernet_list{
    ifnum=1
    ip{
      dhcp=1
      address_list={
        ipv4_address="10.0.2.2"
        ipv4_netmask="255.255.255.0"
      }
    }
    action="enable"
  }
  ethernet_list{
    ifnum=2
    ip{
      dhcp=1
      address_list={
        ipv4_address="10.0.3.10"
        ipv4_netmask="255.255.255.0"
      }
    }
    action="enable"
  }
}
```

## Argument Reference

* `ifnum` - Index of the ethernet interface
* `dhcp` - flag to indicate if dhcp is enabled or disabled
* `ipv4_address` - ipv4 address to be assigned to the ethernet interface
* `ipv4_netmask` - ipv4 subnet mask to be assigned to the ethernet interface
* `action` - flag to indicate if interface is enabled or disabled
