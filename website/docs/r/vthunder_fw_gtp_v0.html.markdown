---
layout: "vthunder"
page_title: "vthunder: vthunder_fw_gtp_v0"
sidebar_current: "docs-vthunder-resource-fw-gtp-v0"
description: |-
	Provides details about vthunder fw gtp v0 resource for A10
---

# vthunder\_fw\_gtp\_v0

`vthunder_fw_gtp_v0` Provides details about vthunder fw gtp v0
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_fw_gtp_v0" "FwTest" {
	gtpv0_value = "enable" 
}
```

## Argument Reference

* `gtpv0_value` - ‘enable’: Enable GTP v0 Inspection;
* `uuid` - uuid of the object

