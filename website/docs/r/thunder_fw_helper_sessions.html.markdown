---
layout: "thunder"
page_title: "thunder: thunder_fw_helper_sessions"
sidebar_current: "docs-thunder-resource-fw-helper-sessions"
description: |-
	Provides details about thunder fw helper sessions resource for A10
---

# thunder\_fw\_helper\_sessions

`thunder_fw_helper_sessions` Provides details about thunder fw helper sessions
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_fw_helper_sessions" "Fw_Helper_Sessions_Test" {

idle_timeout = 0
limit = 0
mode = "string"
uuid = "string"
 
}
```

## Argument Reference

* `idle_timeout` - helper-sessions idle-timeout time (Idle-timeout in minutes (default: 1 minute))
* `limit` - Limit number of helper-sessions (Limit helper-sessions number)
* `mode` - ‘disable’: Disable helper-sessions;
* `uuid` - uuid of the object

