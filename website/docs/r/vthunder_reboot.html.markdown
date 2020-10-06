---
layout: "vthunder"
page_title: "vthunder: vthunder_reboot"
sidebar_current: "docs-vthunder-resource-reboot"
description: |-
    Provides details about vthunder reboot resource for A10
---

# vthunder\_import

`vthunder_reboot` provides details about reboot
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_reboot" "reboot1" {
		all=1 
}
```

## Argument Reference

* `all` - Reboot all devices when VCS is enabled, or only this device itself if VCS is not enabled

