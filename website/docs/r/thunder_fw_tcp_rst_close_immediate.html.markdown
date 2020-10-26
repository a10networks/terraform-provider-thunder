---
layout: "thunder"
page_title: "thunder: thunder_fw_tcp_rst_close_immediate"
sidebar_current: "docs-thunder-resource-fw-tcp-rst-close-immediate"
description: |-
	Provides details about thunder fw tcp rst close immediate resource for A10
---

# thunder\_fw\_tcp\_rst\_close\_immediate

`thunder_fw_tcp_rst_close_immediate` Provides details about thunder fw tcp rst close immediate
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_fw_tcp_rst_close_immediate" "FwTest" {
	status = "enable" 
}
```

## Argument Reference

* `status` - ‘enable’: Enable TCP RST close immediate (default); ‘disable’: Disable TCP RST close immediate;
* `uuid` - uuid of the object

