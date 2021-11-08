---
layout: "thunder"
page_title: "thunder: thunder_fw_local_log"
sidebar_current: "docs-thunder-resource-fw-local-log"
description: |-
    Provides details about thunder fw local log resource for A10
---

# thunder\_fw\_local\_log

`thunder_fw_local_log` Provides details about thunder fw local log
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_fw_local_log" "resourceFwLocalLogTest" {
	local_logging = 0
uuid = "string"
 
}

```

## Argument Reference

* `local-log` - Enable local-log for Application Firewall
* `local-logging` - Enable local logging
* `uuid` - uuid of the object

