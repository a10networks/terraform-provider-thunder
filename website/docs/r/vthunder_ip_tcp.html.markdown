---
layout: "vthunder"
page_title: "vthunder: vthunder_ip_tcp"
sidebar_current: "docs-vthunder-resource-ip-tcp"
description: |-
  Provides details about vthunder ip tcp resource for A10
---

# vthunder\_ip\_tcp

`vthunder_ip_tcp` Provides details about vthunder ip tcp
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_ip_tcp" "Iptcp" {
  syn_cookie  {
      threshold = 4
  }
}
```

## Argument Reference

* `uuid` - uuid of the object
* `threshold` - SYN cookie expire threshold (seconds (default is 4))

