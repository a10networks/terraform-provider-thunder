---
layout: "vthunder"
page_title: "vthunder: vthunder_server"
sidebar_current: "docs-vthunder-resource-server"
description: |-
    Provides details about vthunder server resource for A10
---

# vthunder\_server

`vthunder_server` provides details about configuring backend server on device
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_server" "rs9" {
  health_check_disable=1
  name="rs9"
  host="10.0.3.2"
  port_list {
    health_check_disable=1
    port_number=8080
    protocol="tcp"
  }
}
```

## Argument Reference

* `health_check_disable` - Disable configured health check configuration
* `name` - Server name
* `host` - Server host IP Address
* `port_number` - Port number
* `protocol` - protocol
