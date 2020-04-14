---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_rate_limit_log"
sidebar_current: "docs-vthunder-resource-slb-rate-limit-log"
description: |-
    Provides details about vthunder SLB rate limit log resource for A10
---

# vthunder\_slb\_rate\_limit\_log

`vthunder_slb_rate_limit_log` Provides details about vthunder SLB rate limit log
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_rate_limit_log" "rate_limit" {

	sampling_enable {
		counters1 = "all"
	} 
}
```

## Argument Reference

* `sampling_enable`
    * `counters1` - 'all’: all; 'total_log_times’: Total log times; 'total_log_msg’: Total log messages; 'local_log_msg’: Local log messages; 'remote_log_msg’: Remote log messages; 'local_log_rate’: Local rate (per sec); 'remote_log_rate’: Remote rate (per sec); 'msg_too_big’: Log message too big; 'buff_alloc_fail’: Buffer alloc fail; 'no_route’: No route; 'buff_send_fail’: Buffer send fail; 'alloc_conn’: Log-session alloc; 'free_conn’: Log-session free; 'conn_alloc_fail’: Log-session alloc fail; 'no_repeat_msg’: No repeat message; 'local_log_dropped’: Local log dropped due to rate-limit;