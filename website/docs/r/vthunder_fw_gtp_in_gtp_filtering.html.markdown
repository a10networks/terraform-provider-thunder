---
layout: "vthunder"
page_title: "vthunder: vthunder_fw_gtp_in_gtp_filtering"
sidebar_current: "docs-vthunder-resource-fw-gtp-in-gtp-filtering"
description: |-
	Provides details about vthunder fw gtp in gtp filtering resource for A10
---

# vthunder\_fw\_gtp\_in\_gtp\_filtering

`vthunder_fw_gtp_in_gtp_filtering` Provides details about vthunder fw gtp in gtp filtering
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

* `gtp_in_gtp_value` - ‘disable’: Disable GTP in GTP filtering, (default:Enabled);
* `uuid` - uuid of the object

