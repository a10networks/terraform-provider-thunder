---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_passthrough"
sidebar_current: "docs-vthunder-resource-slb-passthrough"
description: |-
    Provides details about vthunder SLB passthrough resource for A10
---

# vthunder\_slb\_passthrough

`vthunder_slb_passthrough` Provides details about vthunder SLB passthrough
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_passthrough" "testname" {
	sampling_enable  {
	    counters1 = "all"
	}
}
```

## Argument Reference

* `sampling_enable`
    * `counters1` - 'all’: all; 'curr_conn’: Current connections; 'total_conn’: Total connections; 'total_fwd_bytes’: Forward bytes; 'total_fwd_packets’: Forward packets; 'total_rev_bytes’: Reverse bytes; 'total_rev_packets’: Reverse packets; 'curr_pconn’: Persistent connections;