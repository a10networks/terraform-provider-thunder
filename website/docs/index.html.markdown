---
layout: "vthunder"
page_title: "vThunder Provider : Index"
sidebar_current: "docs-vthunder-index"
description: |-
    Provides details about provider vthunder
---

# A10 vThunder Provider

A [Terraform](https://terraform.io) provider for A10 vThunder. Resources are currently available for LTM.

### Requirements

This provider uses the axAPI.

## Example

```
provider "vthunder" {
  address = "${var.url}"
  username = "${var.username}"
  password = "${var.password}"
}
```

## Reference

- `address` - (Required) Address of the device
- `username` - (Required) Username for authentication
- `password` - (Required) Password for authentication
