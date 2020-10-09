---
layout: "thunder"
page_title: "thunder: thunder_fw_urpf"
sidebar_current: "docs-thunder-resource-fw-urpf"
description: |-
	Provides details about thunder fw urpf resource for A10
---

# thunder\_fw\_urpf

`thunder_fw_urpf` Provides details about thunder fw urpf
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_fw_urpf" "FwTest" {
	status = "strict" 
}
```

## Argument Reference

* `status` - ‘disabled’: Disable urpf (Default: disabled); ‘strict’: Perform Strict Check; ‘loose’: Perform Loose Check;
* `uuid` - uuid of the object

