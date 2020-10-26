---
layout: "thunder"
page_title: "thunder: thunder_slb_sport_rate_limit"
sidebar_current: "docs-thunder-slb-sport-rate-limit"
description: |-
    Provides details about thunder SLB sport rate limit resource for A10
---

# thunder\_slb\_sport\_rate\_limit

`thunder_slb_sport_rate_limit` Provides details about thunder SLB sport rate limit
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_sport_rate_limit" "sport_rate" {
	sampling_enable  {
	    counters1 = "all"
	}
}
```

## Argument Reference

* `sampling_enable`
    * `counters1` - 'all’: all; 'alloc_sport’: Alloc’d src port entry; 'alloc_sportip’: Alloc’d src port-ip entry; 'freed_sport’: Freed src port entry; 'freed_sportip’: Freed src port-ip entry; 'total_drop’: Total rate exceed drop; 'total_reset’: Total rate exceed reset; 'total_log’: Total log sent;




