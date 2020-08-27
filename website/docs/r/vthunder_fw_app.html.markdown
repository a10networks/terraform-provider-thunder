---
layout: "vthunder"
page_title: "vthunder: vthunder_fw_app"
sidebar_current: "docs-vthunder-resource-fw-app"
description: |-
	Provides details about vthunder fw app resource for A10
---

# vthunder\_fw\_app

`vthunder_fw_app` Provides details about vthunder fw app
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

* `uuid` - uuid of the object
* `counters1` - ‘all’: all; ‘dummy’: Entry for a10countergen;

