---
layout: "thunder"
page_title: "thunder: thunder_reboot"
sidebar_current: "docs-thunder-resource-reboot"
description: |-
    Provides details about thunder reboot resource for A10
---

# thunder\_import

`thunder_reboot` provides details about reboot
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_reboot" "reboot1" {
		all=1 
}
```

## Argument Reference

* `all` - Reboot all devices when VCS is enabled, or only this device itself if VCS is not enabled

