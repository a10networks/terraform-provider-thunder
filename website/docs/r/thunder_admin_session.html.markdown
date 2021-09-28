---
layout: "thunder"
page_title: "thunder: thunder_admin_session"
sidebar_current: "docs-thunder-resource-admin-session"
description: |-
    Provides details about thunder admin session resource for A10
---

# thunder\_admin\_session

`thunder_admin_session` Provides details about thunder admin session
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_admin_session" "resourceAdminSessionTest" {
	clear = 0
all = 0
sid = 0
uuid = "string"
 
}

```

## Argument Reference

* `admin-session` - Admin session management
* `clear` - clear admin session
* `all` - Clear all admin sessions
* `sid` - Session ID
* `uuid` - uuid of the object

