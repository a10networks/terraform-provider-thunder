---
layout: "vthunder"
page_title: "vthunder: vthunder_fw_local_log"
sidebar_current: "docs-vthunder-resource-fw-local-log"
description: |-
	Provides details about vthunder fw local log resource for A10
---

# vthunder\_fw\_local\_log

`vthunder_fw_local_log` Provides details about vthunder fw local log
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

* `local_logging` - Enable local logging
* `uuid` - uuid of the object

