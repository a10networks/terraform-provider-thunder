---
layout: "thunder"
page_title: "thunder: thunder_fw_gtp_in_gtp_filtering"
sidebar_current: "docs-thunder-resource-fw-gtp-in-gtp-filtering"
description: |-
	Provides details about thunder fw gtp in gtp filtering resource for A10
---

# thunder\_fw\_gtp\_in\_gtp\_filtering

`thunder_fw_gtp_in_gtp_filtering` Provides details about thunder fw gtp in gtp filtering
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_fw_gtp_in_gtp_filtering" "FwTest" {
	gtp_in_gtp_value = "disable" 
}
```

## Argument Reference

* `gtp_in_gtp_value` - ‘disable’: Disable GTP in GTP filtering, (default:Enabled);
* `uuid` - uuid of the object

