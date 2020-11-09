---
layout: "thunder"
page_title: "thunder: thunder_fw_tcp_reset_on_error"
sidebar_current: "docs-thunder-resource-fw-tcp-reset-on-error"
description: |-
	Provides details about thunder fw tcp reset on error resource for A10
---

# thunder\_fw\_tcp\_reset\_on\_error

`thunder_fw_tcp_reset_on_error` Provides details about thunder fw tcp reset on error
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_fw_tcp_reset_on_error" "Fw_Tcp_Reset_On_Error_Test" {
enable = 0
uuid = "string"
 
}
```

## Argument Reference

* `enable` - Enable send TCP reset on error
* `uuid` - uuid of the object
