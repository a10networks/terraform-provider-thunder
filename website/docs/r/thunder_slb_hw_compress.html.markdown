---
layout: "thunder"
page_title: "thunder: thunder_slb_hw_compress"
sidebar_current: "docs-thunder-resource-slb-hw-compress"
description: |-
    Provides details about thunder SLB hw-compress resource for A10
---

# thunder\_slb\_hw_compress

`thunder_slb_hw_compress` Provides details about thunder SLB hw-compress
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_hw_compress" "hw_compress" {
	sampling_enable {
		counters1 = "all"
	}
}
```

## Argument Reference

* `sampling_enable` - 
	* `counters1` - 'all’: all; 'request_count’: Total request count; 'submit_count’: Total submit count; 'response_count’: Total response count; 'failure_count’: Total failure count; 'failure_code’: Last failure code; 'ring_full_count’: Compression queue full; 'max_outstanding_request_count’: Max queued request count; 'max_outstanding_submit_count’: Max queued submit count;


