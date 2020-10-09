---
layout: "thunder"
page_title: "thunder: thunder_vrrp_common"
sidebar_current: "docs-thunder-resource-vrrp-common"
description: |-
    Provides details about thunder vrrp common resource for A10
---

# thunder\_vrrp\_common

`thunder_vrrp_common` provides details about vrrp common
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

	resource "thunder_vrrp_common" "vrrp_common" {
		set_id = 1
		device_id = 1
		action = "enable"
}
```

## Argument Reference

* `set_id` - Set-ID for HA configuration (Set id from 1 to 15)
* `device_id` - Unique ID for each VRRP-A box (Device-id number)
* `action` - 'enable'= enable vrrp-a; 'disable'= disable vrrp-a; 

