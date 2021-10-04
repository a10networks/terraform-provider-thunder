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
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_slb_template_server" "resourceSlbTemplateServerTest" {
	name = "string"
conn_limit = 0
resume = 0
conn_limit_no_logging = 0
conn_rate_limit = 0
rate_interval = "string"
conn_rate_limit_no_logging = 0
dns_query_interval = 0
dns_fail_interval = 0
dynamic_server_prefix = "string"
extended_stats = 0
log_selection_failure = 0
health_check = "string"
health_check_disable = 0
max_dynamic_server = 0
min_ttl_ratio = 0
weight = 0
spoofing_cache = 0
stats_data_action = "string"
slow_start = 0
initial_slow_start = 0
add = 0
times = 0
every = 0
till = 0
bw_rate_limit_acct = "string"
bw_rate_limit = 0
bw_rate_limit_resume = 0
bw_rate_limit_duration = 0
bw_rate_limit_no_logging = 0
uuid = "string"
user_tag = "string"
 
}

```

## Argument Reference

* `server` - Server template
* `name` - Server template name
* `conn-limit` - Connection limit
* `resume` - Resume accepting new connection after connection number drops below threshold (Connection resume threshold)
* `conn-limit-no-logging` - Do not log connection over limit event
* `conn-rate-limit` - Connection rate limit
* `rate-interval` - '100ms': Use 100 ms as sampling interval; 'second': Use 1 second as sampling interval;
* `conn-rate-limit-no-logging` - Do not log connection over limit event
* `dns-query-interval` - The interval to query DNS server for the hostname (DNS query interval (in minute, default is 10))
* `dns-fail-interval` - The interval to retry when DNS failed to query (DNS failure interval (in second, default is 30))
* `dynamic-server-prefix` - Prefix of dynamic server (Prefix of dynamic server (default is "DRS"))
* `extended-stats` - Enable extended statistics on real server
* `log-selection-failure` - Enable real-time logging for server selection failure event
* `health-check` - Health Check Monitor (Health monitor name)
* `health-check-disable` - Disable configured health check configuration
* `max-dynamic-server` - Maximum dynamic server number (Maximum dynamic server number (default is 255))
* `min-ttl-ratio` - Minimum TTL to DNS query interval ratio (Minimum TTL ratio (default is 2))
* `weight` - Weight for the Real Servers (Connection Weight (default is 1))
* `spoofing-cache` - Servers under the template are spoofing cache
* `stats-data-action` - 'stats-data-enable': Enable statistical data collection for real server; 'stats-data-disable': Disable statistical data collection for real server;
* `slow-start` - Slowly ramp up the connection number after server is up
* `initial-slow-start` - Initial slow start connection limit (default 128)
* `add` - Slow start connection limit add by a number every interval (Add by this number every interval)
* `times` - Slow start connection limit multiply by a number every interval (default 2) (Multiply by this number every interval)
* `every` - Slow start connection limit increment interval (default 10)
* `till` - Slow start ends when slow start connection limit reaches a number (default 4096) (Slow start ends when connection limit reaches this number)
* `bw-rate-limit-acct` - 'to-server-only': Only account for traffic sent to server; 'from-server-only': Only account for traffic received from server; 'all': Account for all traffic sent to and received from server;
* `bw-rate-limit` - Configure bandwidth rate limit on real server (Bandwidth rate limit in Kbps)
* `bw-rate-limit-resume` - Resume server selection after bandwidth drops below this threshold (in Kbps) (Bandwidth rate limit resume threshold (in Kbps))
* `bw-rate-limit-duration` - Duration in seconds the observed rate needs to honor
* `bw-rate-limit-no-logging` - Do not log bandwidth rate limit related state transitions
* `uuid` - uuid of the object
* `user-tag` - Customized tag

