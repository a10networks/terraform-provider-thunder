---
layout: "thunder"
page_title: "thunder: thunder_service_group"
sidebar_current: "docs-thunder-resource-service-group"
description: |-
    Provides details about thunder service group resource for A10
---

# thunder\_service\_group

`thunder_service_group` provides details about configuring service group on device
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_service_group" "sg9" {
  name="sg9"
  protocol="TCP"
  member_list {
    name="rs9"
    port=8080
  }
}
```

## Argument Reference

* `name` - Name of service group
* `protocol` - protocol
* `port` - Port number
