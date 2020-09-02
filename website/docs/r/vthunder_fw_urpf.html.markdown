---
layout: "vthunder"
page_title: "vthunder: vthunder_fw_urpf"
sidebar_current: "docs-vthunder-resource-fw-urpf"
description: |-
	Provides details about vthunder fw urpf resource for A10
---

# vthunder\_fw\_urpf

`vthunder_fw_urpf` Provides details about vthunder fw urpf
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_fw_urpf" "FwTest" {
	status = "strict" 
}
```

## Argument Reference

* `status` - ‘disabled’: Disable urpf (Default: disabled); ‘strict’: Perform Strict Check; ‘loose’: Perform Loose Check;
* `uuid` - uuid of the object

