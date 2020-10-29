---
layout: "thunder"
page_title: "thunder: thunder_slb_template_server"
sidebar_current: "docs-thunder-resource-slb-template-server"
description: |-
	Provides details about thunder slb template server resource for A10
---

# thunder\_slb\_template\_server

`thunder_slb_template_server` Provides details about thunder slb template server
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_server" "Slb_Template_Server_Test" {

health_check_disable = 0
stats_data_action = "string"
slow_start = 0
weight = 0
bw_rate_limit = 0
spoofing_cache = 0
conn_limit = 0
uuid = "string"
resume = 0
max_dynamic_server = 0
rate_interval = "string"
till = 0
add = 0
min_ttl_ratio = 0
bw_rate_limit_no_logging = 0
dynamic_server_prefix = "string"
initial_slow_start = 0
every = 0
conn_limit_no_logging = 0
extended_stats = 0
conn_rate_limit_no_logging = 0
name = "string"
bw_rate_limit_duration = 0
bw_rate_limit_resume = 0
bw_rate_limit_acct = "string"
user_tag = "string"
times = 0
log_selection_failure = 0
conn_rate_limit = 0
dns_query_interval = 0
health_check = "string"
 
}
```

## Argument Reference

* `add` - Slow start connection limit add by a number every interval (Add by this number every interval)
* `bw_rate_limit` - Configure bandwidth rate limit on real server (Bandwidth rate limit in Kbps)
* `bw_rate_limit_acct` - ‘to-server-only’: Only account for traffic sent to server; ‘from-server-only’: Only account for traffic received from server; ‘all’: Account for all traffic sent to and received from server;
* `bw_rate_limit_duration` - Duration in seconds the observed rate needs to honor
* `bw_rate_limit_no_logging` - Do not log bandwidth rate limit related state transitions
* `bw_rate_limit_resume` - Resume server selection after bandwidth drops below this threshold (in Kbps) (Bandwidth rate limit resume threshold (in Kbps))
* `conn_limit` - Connection limit
* `conn_limit_no_logging` - Do not log connection over limit event
* `conn_rate_limit` - Connection rate limit
* `conn_rate_limit_no_logging` - Do not log connection over limit event
* `dns_query_interval` - The interval to query DNS server for the hostname (DNS query interval (in minute, default is 10))
* `dynamic_server_prefix` - Prefix of dynamic server (Prefix of dynamic server (default is “DRS”))
* `every` - Slow start connection limit increment interval (default 10)
* `extended_stats` - Enable extended statistics on real server
* `health_check` - Health Check Monitor (Health monitor name)
* `health_check_disable` - Disable configured health check configuration
* `initial_slow_start` - Initial slow start connection limit (default 128)
* `log_selection_failure` - Enable real-time logging for server selection failure event
* `max_dynamic_server` - Maximum dynamic server number (Maximum dynamic server number (default is 255))
* `min_ttl_ratio` - Minimum TTL to DNS query interval ratio (Minimum TTL ratio (default is 2))
* `name` - Server template name
* `rate_interval` - ‘100ms’: Use 100 ms as sampling interval; ‘second’: Use 1 second as sampling interval;
* `resume` - Resume accepting new connection after connection number drops below threshold (Connection resume threshold)
* `slow_start` - Slowly ramp up the connection number after server is up
* `spoofing_cache` - Servers under the template are spoofing cache
* `stats_data_action` - ‘stats-data-enable’: Enable statistical data collection for real server; ‘stats-data-disable’: Disable statistical data collection for real server;
* `till` - Slow start ends when slow start connection limit reaches a number (default 4096) (Slow start ends when connection limit reaches this number)
* `times` - Slow start connection limit multiply by a number every interval (default 2) (Multiply by this number every interval)
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `weight` - Weight for the Real Servers (Connection Weight)
