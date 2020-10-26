---
layout: "thunder"
page_title: "thunder: thunder_vrrp_session_sync"
sidebar_current: "docs-thunder-resource-vrrp-session-sync"
description: |-
    Provides details about thunder vrrp session sync resource for A10
---

# thunder\_vrrp\_session\_sync

`thunder_vrrp_session_sync` provides details about vrrp session sync
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_vrrp_session_sync" "sync" {
		action="enable"
}
```

## Argument Reference

* `action` - 'enable'= enable vrrp-a session sync (default); 'disable'= disable vrrp-a session sync; "


