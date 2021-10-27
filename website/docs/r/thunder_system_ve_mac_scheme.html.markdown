---
layout: "thunder"
page_title: "thunder: thunder_system_ve_mac_scheme"
sidebar_current: "docs-thunder-resource-system-ve-mac-scheme"
description: |-
    Provides details about thunder system ve mac scheme resource for A10
---

# thunder\_system\_ve\_mac\_scheme

`thunder_system_ve_mac_scheme` Provides details about thunder system ve mac scheme
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_system_ve_mac_scheme" "resourceSystemVeMacSchemeTest" {
	ve_mac_scheme_val = "string"
uuid = "string"
 
}

```

## Argument Reference

* `ve-mac-scheme` - VE MAC allocation scheme
* `ve-mac-scheme-val` - 'hash-based': Hash-based using the VE number; 'round-robin': Round Robin scheme; 'system-mac': Use system MAC address;
* `uuid` - uuid of the object

