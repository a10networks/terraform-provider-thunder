---
layout: "thunder"
page_title: "Thunder Provider : Index"
sidebar_current: "docs-thunder-index"
description: |-
    Provides details about provider thunder
---

# A10 Thunder Provider

A [Terraform](https://terraform.io) provider for A10 Thunder.

### Requirements

This provider uses the aXAPI.

## Example

```
provider "thunder" {
  address    = "${var.url}"
  username   = "${var.username}"
  password   = "${var.password}"
  partition  = "${var.partition} 
}
```

## Reference

- `address` - (Required) Address of the device
- `username` - (Required) Username for authentication
- `password` - (Required) Password for authentication
- `partition` - (Optional) If any specific partition to use; default is "shared"