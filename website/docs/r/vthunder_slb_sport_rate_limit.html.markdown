---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_sport_rate_limit"
sidebar_current: "docs-vthunder-slb-sport-rate-limit"
description: |-
    Provides details about vthunder SLB sport rate limit resource for A10
---

# vthunder\_slb\_sport\_rate\_limit

`vthunder_slb_sport_rate_limit` Provides details about vthunder SLB sport rate limit
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_sport_rate_limit" "testname" {
	sampling_enable  {
	    counters1 = "all"
	}
}
```

## Argument Reference

* `sampling_enable`
    * `counters1` - 'all’: all; 'alloc_sport’: Alloc’d src port entry; 'alloc_sportip’: Alloc’d src port-ip entry; 'freed_sport’: Freed src port entry; 'freed_sportip’: Freed src port-ip entry; 'total_drop’: Total rate exceed drop; 'total_reset’: Total rate exceed reset; 'total_log’: Total log sent;




