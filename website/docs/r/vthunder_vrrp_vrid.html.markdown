---
layout: "vthunder"
page_title: "vthunder: vthunder_vrrp_vrid"
sidebar_current: "docs-vthunder-resource-vrrp-vrid"
description: |-
    Provides details about vthunder vrrp vrid resource for A10
---

# vthunder\_vrrp\_vrid

`vthunder_vrrp_vrid provides details about vrrp vrid
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_vrrp_vrid" "vrrp_vrid" {
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


