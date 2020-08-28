---
layout: "vthunder"
page_title: "vthunder: vthunder_fw_clear_session_filter"
sidebar_current: "docs-vthunder-resource-fw-clear-session-filter"
description: |-
	Provides details about vthunder fw clear session filter resource for A10
---

# vthunder\_fw\_clear\_session\_filter

`vthunder_fw_clear_session_filter` Provides details about vthunder fw clear session filter
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

* `status` - ‘disable’: Disable clear L4 session filter for fw (Default: disabled); ‘enable’: Enable clear L4 session filter for fw;
* `uuid` - uuid of the object

