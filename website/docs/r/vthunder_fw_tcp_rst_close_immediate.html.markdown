---
layout: "vthunder"
page_title: "vthunder: vthunder_fw_tcp_rst_close_immediate"
sidebar_current: "docs-vthunder-resource-fw-tcp-rst-close-immediate"
description: |-
	Provides details about vthunder fw tcp rst close immediate resource for A10
---

# vthunder\_fw\_tcp\_rst\_close\_immediate

`vthunder_fw_tcp_rst_close_immediate` Provides details about vthunder fw tcp rst close immediate
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

* `status` - ‘enable’: Enable TCP RST close immediate (default); ‘disable’: Disable TCP RST close immediate;
* `uuid` - uuid of the object

