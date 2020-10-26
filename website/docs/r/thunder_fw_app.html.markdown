---
layout: "thunder"
page_title: "thunder: thunder_fw_app"
sidebar_current: "docs-thunder-resource-fw-app"
description: |-
	Provides details about thunder fw app resource for A10
---

# thunder\_fw\_app

`thunder_fw_app` Provides details about thunder fw app
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_fw_app" "FwTest" {
	sampling_enable {
		counters1 = "all" 
	}
}
```

## Argument Reference

* `uuid` - uuid of the object
* `counters1` - ‘all’: all; ‘dummy’: Entry for a10countergen;

