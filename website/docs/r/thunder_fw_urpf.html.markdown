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
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_fw_urpf" "resourceFwUrpfTest" {
	status = "string"
uuid = "string"
 
}

```

## Argument Reference

* `urpf` - Unicast Reverse Path Forwarding (Default: loose)
* `status` - 'loose': Perform loose check; 'strict': Perform strict check; 'disable': Disable check;
* `uuid` - uuid of the object

