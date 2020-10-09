---
layout: "thunder"
page_title: "thunder: thunder_slb_passthrough"
sidebar_current: "docs-thunder-resource-slb-passthrough"
description: |-
    Provides details about thunder SLB passthrough resource for A10
---

# thunder\_slb\_passthrough

`thunder_slb_passthrough` Provides details about thunder SLB passthrough
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_passthrough" "passthrough" {
	sampling_enable  {
	    counters1 = "all"
	}
}
```

## Argument Reference

* `sampling_enable`
    * `counters1` - 'all’: all; 'curr_conn’: Current connections; 'total_conn’: Total connections; 'total_fwd_bytes’: Forward bytes; 'total_fwd_packets’: Forward packets; 'total_rev_bytes’: Reverse bytes; 'total_rev_packets’: Reverse packets; 'curr_pconn’: Persistent connections;