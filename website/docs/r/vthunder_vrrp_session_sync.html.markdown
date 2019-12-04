---
layout: "vthunder"
page_title: "vthunder: vthunder_vrrp_session_sync"
sidebar_current: "docs-vthunder-resource-vrrp-session-sync"
description: |-
    Provides details about vthunder vrrp session sync resource for A10
---

# vthunder\_vrrp\_session\_sync

`vthunder_vrrp_session_sync provides details about vrrp session sync
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_vrrp_session_sync" "sync" {
		action="enable"
}
```

## Argument Reference

* `action` - 'enable'= enable vrrp-a session sync (default); 'disable'= disable vrrp-a session sync; "


