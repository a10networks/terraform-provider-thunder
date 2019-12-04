---
layout: "vthunder"
page_title: "vthunder: vthunder_vrrp_common"
sidebar_current: "docs-vthunder-resource-vrrp-common"
description: |-
    Provides details about vthunder vrrp common resource for A10
---

# vthunder\_vrrp\_common

`vthunder_vrrp_common provides details about vrrp common
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

	resource "vthunder_vrrp_common" "vrrp_common" {
		set_id = 1
		device_id = 1
		action = "enable"
}
```

## Argument Reference

* `set_id` - Set-ID for HA configuration (Set id from 1 to 15)
* `device_id` - Unique ID for each VRRP-A box (Device-id number)
* `action` - 'enable'= enable vrrp-a; 'disable'= disable vrrp-a; 

