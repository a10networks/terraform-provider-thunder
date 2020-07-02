---
layout: "vthunder"
page_title: "vthunder: vthunder_ip_address"
sidebar_current: "docs-vthunder-resource-ip-address"
description: |-
	Provides details about vthunder ip address resource for A10
---

# vthunder\_ip\_address

`vthunder_ip_address` Provides details about vthunder ip address
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

// Put working JSON here
```

## Argument Reference

* `ip_addr` - IP address
* `ip_mask` - IP subnet mask
* `uuid` - uuid of the object

