---
layout: "vthunder"
page_title: "vthunder: vthunder_service_group"
sidebar_current: "docs-vthunder-resource-service-group"
description: |-
    Provides details about vthunder service group resource for A10
---

# vthunder\_service\_group

`vthunder_service_group` provides details about configuring service group on device
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_service_group" "sg9" {
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
