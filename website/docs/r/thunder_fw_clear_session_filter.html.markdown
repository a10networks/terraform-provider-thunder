---
layout: "thunder"
page_title: "thunder: thunder_fw_clear_session_filter"
sidebar_current: "docs-thunder-resource-fw-clear-session-filter"
description: |-
	Provides details about thunder fw clear session filter resource for A10
---

# thunder\_fw\_clear\_session\_filter

`thunder_fw_clear_session_filter` Provides details about thunder fw clear session filter
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_fw_clear_session_filter" "Fw_Clear_Session_Filter_Test" {
status = "string"
uuid = "string"
 
}

```

## Argument Reference

* `status` - ‘disable’: Disable clear L4 session filter for fw (Default: disabled); ‘enable’: Enable clear L4 session filter for fw;
* `uuid` - uuid of the object
