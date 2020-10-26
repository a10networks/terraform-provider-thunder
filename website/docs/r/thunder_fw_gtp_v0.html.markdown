---
layout: "thunder"
page_title: "thunder: thunder_fw_gtp_v0"
sidebar_current: "docs-thunder-resource-fw-gtp-v0"
description: |-
	Provides details about thunder fw gtp v0 resource for A10
---

# thunder\_fw\_gtp\_v0

`thunder_fw_gtp_v0` Provides details about thunder fw gtp v0
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_fw_gtp_v0" "FwTest" {
	gtpv0_value = "enable" 
}
```

## Argument Reference

* `gtpv0_value` - ‘enable’: Enable GTP v0 Inspection;
* `uuid` - uuid of the object

