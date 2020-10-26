---
layout: "thunder"
page_title: "thunder: thunder_slb_common"
sidebar_current: "docs-thunder-resource-slb-common"
description: |-
    Provides details about thunder SLB common resource for A10
---

# thunder\_slb\_common

`thunder_slb_common` Provides details about thunder SLB common
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_common" "common" {
	low_latency = 0
	use_mss_tab = 0
	stats_data_disable = 0
	compress_block_size = 7000
	player_id_check_enable = 0
	msl_time = 20
	graceful_shutdown_enable = 0
	hw_syn_rr = 2000 
	conn_rate_limit {
		src_ip_list {
			protocol = "tcp"
			limit_period = "100"
			limit = 50
			exceed_action = 0
			shared = 0
		}
	}
}
```

## Argument Reference

* `low_latency` - Enable low latency mode
* `use_mss_tab` - Use MSS based on internal table for SLB processing
* `stats_data_disable` - Disable global slb data statistics
* `compress_block_size` - Set compression block size (Compression block size in bytes)
* `player_id_check_enable` - Enable the Player id check
* `msl_time` - Configure maximum session life, default is 2 seconds (1-40 seconds, default is 2 seconds)
* `graceful_shutdown_enable` - Enable graceful shutdown
* `hw_syn_rr` - Configure hardware SYN round robin (range 1-500000)
* `conn_rate_limit` - 
    * `src_ip_list` - 
        * `protocol` - 'tcp’: Set TCP connection rate limit; 'udp’: Set UDP packet rate limit;
        * `limit_period` - '100’: 100 ms; '1000’: 1000 ms;
        * `limit` - Set max connections per period
        * `exceed_action` - Set action if threshold exceeded
        * `shared` - Set threshold shared amongst all virtual ports
