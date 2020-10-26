---
layout: "thunder"
page_title: "thunder: thunder_fw_tcp_mss_clamp"
sidebar_current: "docs-thunder-resource-fw-tcp-mss-clamp"
description: |-
	Provides details about thunder fw tcp mss clamp resource for A10
---

# thunder\_fw\_tcp\_mss\_clamp

`thunder_fw_tcp_mss_clamp` Provides details about thunder fw tcp mss clamp
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_fw_tcp_mss_clamp" "FwTcpTest" {
	mss_clamp_type = "fixed"
	mss_value = "0" 
}
```

## Argument Reference

* `min` - Specify the min value allowed for the TCP MSS (Specify the min value allowed for the TCP MSS (default: ((576 - 60 - 60))))
* `mss_clamp_type` - ‘fixed’: Specify a fixed max value for the TCP MSS; ‘subtract’: Specify the value to subtract from the TCP MSS;
* `mss_subtract` - Specify the value to subtract from the TCP MSS (default: not configured)
* `mss_value` - The max value allowed for the TCP MSS (default: not configured)}
* `uuid` - uuid of the object

