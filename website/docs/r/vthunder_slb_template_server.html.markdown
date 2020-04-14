---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_server"
sidebar_current: "docs-vthunder-resource-slb_template_server"
description: |-
    Provides details about vthunder SLB template server resource for A10
---

# vthunder\_slb\_template\_server

`vthunder_slb_template_server` Provides details about vthunder SLB template server
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_server" "server" {
	name = "testserver"
	user_tag = "test_tag"
	stats_data_action = "stats-data-enable"
	slow_start = 1
	weight = 3
	bw_rate_limit = 2
	spoofing_cache = 1
	conn_limit = 1
	resume = 1
	max_dynamic_server = 3
	rate_interval = "second"
	min_ttl_ratio = 2
	bw_rate_limit_no_logging = 1
	dynamic_server_prefix = "DRS"
	conn_limit_no_logging = 1
	extended_stats = 1
	conn_rate_limit_no_logging = 1
	bw_rate_limit_duration = 1
	bw_rate_limit_resume = 1
	bw_rate_limit_acct = "to-server-only"
	log_selection_failure = 1
	conn_rate_limit = 5
	dns_query_interval = 2
}
```

## Argument Reference

* `name` - Server template name
* `user_tag` - Customized tag
* `stats_data_action` - 'stats-data-enable’: Enable statistical data collection for real server; 'stats-data-disable’: Disable statistical data collection for real server;
* `slow_start` - Slowly ramp up the connection number after server is up
* `weight` - Weight for the Real Servers (Connection Weight (default is 1))
* `bw_rate_limit` - Configure bandwidth rate limit on real server (Bandwidth rate limit in Kbps)
* `spoofing_cache` - Servers under the template are spoofing cache
* `conn_limit` - Connection limit
* `resume` - Resume accepting new connection after connection number drops below threshold (Connection resume threshold)
* `max_dynamic_server` - Maximum dynamic server number (Maximum dynamic server number (default is 255))
* `rate_interval` - '100ms’: Use 100 ms as sampling interval; 'second’: Use 1 second as sampling interval;
* `min_ttl_ratio` - Minimum TTL to DNS query interval ratio (Minimum TTL ratio (default is 2))
* `bw_rate_limit_no_logging` - Do not log bandwidth rate limit related state transitions
* `dynamic_server_prefix` - Prefix of dynamic server (Prefix of dynamic server (default is “DRS”))
* `conn_limit_no_logging` - Do not log connection over limit event
* `extended_stats` - Enable extended statistics on real server
* `conn_rate_limit_no_logging` - Do not log connection over limit event
* `bw_rate_limit_duration` - Duration in seconds the observed rate needs to honor
* `bw_rate_limit_resume` - Resume server selection after bandwidth drops below this threshold (in Kbps) (Bandwidth rate limit resume threshold (in Kbps))
* `bw_rate_limit_acct` - 'to-server-only’: Only account for traffic sent to server; 'from-server-only’: Only account for traffic received from server; 'all’: Account for all traffic sent to and received from server;
* `log_selection_failure` - Enable real-time logging for server selection failure event
* `conn_rate_limit` - Connection rate limit
* `dns_query_interval` - The interval to query DNS server for the hostname (DNS query interval (in minute, default is 10))


