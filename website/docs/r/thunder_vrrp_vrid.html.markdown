---
layout: "thunder"
page_title: "thunder: thunder_vrrp_vrid"
sidebar_current: "docs-thunder-resource-vrrp-vrid"
description: |-
    Provides details about thunder vrrp vrid resource for A10
---

# thunder\_vrrp\_vrid

`thunder_vrrp_vrid` provides details about vrrp vrid
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_vrrp_vrid" "vrrp_vrid" {
	  vrid_val = 1
      floating_ip {
        ip_address_cfg {
        ip_address = "1.1.1.10"
       }
      ip_address_cfg {
      ip_address = "1.1.1.11"
     }
    }
      blade_parameters {
      priority = 250
   }
}
```

## Argument Reference

* `vrid_val` - "Specify ha VRRP-A vrid"
* `ip_address` - IP address field
* `priority` - blade parameter priority


