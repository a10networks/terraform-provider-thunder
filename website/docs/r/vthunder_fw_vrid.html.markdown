---
layout: "vthunder"
page_title: "vthunder: vthunder_fw_vrid"
sidebar_current: "docs-vthunder-resource-fw-vrid"
description: |-
	Provides details about vthunder fw vrid resource for A10
---

# vthunder\_fw\_vrid

`vthunder_fw_vrid` Provides details about vthunder fw vrid
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_fw_vrid" "FwTest" {
	vrid = 1 
}
```

## Argument Reference

* `uuid` - uuid of the object
* `vrid` - Vrrp group (VRRP-A vrid)

