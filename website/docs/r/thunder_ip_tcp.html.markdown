---
layout: "thunder"
page_title: "thunder: thunder_ip_tcp"
sidebar_current: "docs-thunder-resource-ip-tcp"
description: |-
  Provides details about thunder ip tcp resource for A10
---

# thunder\_ip\_tcp

`thunder_ip_tcp` Provides details about thunder ip tcp
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_ip_tcp" "Iptcp" {
  syn_cookie  {
      threshold = 4
  }
}
```

## Argument Reference

* `uuid` - uuid of the object
* `threshold` - SYN cookie expire threshold (seconds (default is 4))

