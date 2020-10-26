---
layout: "thunder"
page_title: "thunder: thunder_server"
sidebar_current: "docs-thunder-resource-server"
description: |-
    Provides details about thunder server resource for A10
---

# thunder\_server

`thunder_server` provides details about configuring backend server on device
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_server" "rs9" {
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
