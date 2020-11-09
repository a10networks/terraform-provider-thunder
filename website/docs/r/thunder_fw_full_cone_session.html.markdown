---
layout: "thunder"
page_title: "thunder: thunder_fw_full_cone_session"
sidebar_current: "docs-thunder-resource-fw-full-cone-session"
description: |-
	Provides details about thunder fw full cone session resource for A10
---

# thunder\_fw\_full\_cone\_session

`thunder_fw_full_cone_session` Provides details about thunder fw full cone session
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_fw_full_cone_session" "Fw_Full_Cone_Session_Test" {
        uuid = "string"
 
}

```

## Argument Reference

* `uuid` - uuid of the object
