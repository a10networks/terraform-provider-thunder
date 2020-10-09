---
layout: "thunder"
page_title: "thunder: thunder_fw_vrid"
sidebar_current: "docs-thunder-resource-fw-vrid"
description: |-
	Provides details about thunder fw vrid resource for A10
---

# thunder\_fw\_vrid

`thunder_fw_vrid` Provides details about thunder fw vrid
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_fw_vrid" "FwTest" {
	vrid = 1 
}
```

## Argument Reference

* `uuid` - uuid of the object
* `vrid` - Vrrp group (VRRP-A vrid)

