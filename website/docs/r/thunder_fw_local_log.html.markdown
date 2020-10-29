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
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_fw_local_log" "Fw_Local_Log_Test" {
local_logging = 0
uuid = "string"
 
}
```

## Argument Reference

* `local_logging` - Enable local logging
* `uuid` - uuid of the object
