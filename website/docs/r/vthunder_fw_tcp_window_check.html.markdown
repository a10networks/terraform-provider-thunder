---
layout: "vthunder"
page_title: "vthunder: vthunder_fw_tcp_window_check"
sidebar_current: "docs-vthunder-resource-fw-tcp-window-check"
description: |-
	Provides details about vthunder fw tcp window check resource for A10
---

# vthunder\_fw\_tcp\_window\_check

`vthunder_fw_tcp_window_check` Provides details about vthunder fw tcp window check
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_fw_tcp_window_check" "FwTest" {
        status = "enable" 
}
```

## Argument Reference

* `status` - ‘enable’: Enable TCP window check (default); ‘disable’: Disable TCP window check;
* `uuid` - uuid of the object
* `counters1` - ‘all’: all; ‘outside-window’: packet dropped for outside of tcp window;

