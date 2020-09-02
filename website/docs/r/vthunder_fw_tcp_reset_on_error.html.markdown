---
layout: "vthunder"
page_title: "vthunder: vthunder_fw_tcp_reset_on_error"
sidebar_current: "docs-vthunder-resource-fw-tcp-reset-on-error"
description: |-
	Provides details about vthunder fw tcp reset on error resource for A10
---

# vthunder\_fw\_tcp\_reset\_on\_error

`vthunder_fw_tcp_reset_on_error` Provides details about vthunder fw tcp reset on error
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

* `enable` - Enable send TCP reset on error
* `uuid` - uuid of the object

