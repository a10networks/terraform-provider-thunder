---
layout: "thunder"
page_title: "thunder: thunder_fw_tcp_window_check"
sidebar_current: "docs-thunder-resource-fw-tcp-window-check"
description: |-
	Provides details about thunder fw tcp window check resource for A10
---

# thunder\_fw\_tcp\_window\_check

`thunder_fw_tcp_window_check` Provides details about thunder fw tcp window check
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_fw_tcp_window_check" "FwTest" {
        status = "enable" 
}
```

## Argument Reference

* `status` - ‘enable’: Enable TCP window check (default); ‘disable’: Disable TCP window check;
* `uuid` - uuid of the object
* `counters1` - ‘all’: all; ‘outside-window’: packet dropped for outside of tcp window;

