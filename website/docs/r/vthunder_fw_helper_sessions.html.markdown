---
layout: "vthunder"
page_title: "vthunder: vthunder_fw_helper_sessions"
sidebar_current: "docs-vthunder-resource-fw-helper-sessions"
description: |-
	Provides details about vthunder fw helper sessions resource for A10
---

# vthunder\_fw\_helper\_sessions

`vthunder_fw_helper_sessions` Provides details about vthunder fw helper sessions
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_fw_helper_sessions" "FwTest" {
	mode = "disable"
}
```

## Argument Reference

* `idle_timeout` - helper-sessions idle-timeout time (Idle-timeout in minutes (default: 1 minute))
* `limit` - Limit number of helper-sessions (Limit helper-sessions number)
* `mode` - ‘disable’: Disable helper-sessions;
* `uuid` - uuid of the object

