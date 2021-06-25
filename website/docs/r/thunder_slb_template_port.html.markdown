---
layout: "thunder"
page_title: "thunder: thunder_slb_template_port"
sidebar_current: "docs-thunder-resource-slb-template-port"
description: |-
	Provides details about thunder slb template port resource for A10
---

# thunder\_slb\_template\_port

`thunder_slb_template_port` Provides details about thunder slb template port
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_slb_template_port" "Slb_Template_Port_Test" {
	
health_check_disable = 0
stats_data_action = "string"
resel_on_reset = 0
dest_nat = 0
restore_svc_time = 0
request_rate_limit = 0
dynamic_member_priority = 0
bw_rate_limit = 0
slow_start = 0
decrement = 0
conn_limit = 0
retry = 0
weight = 0
inband_health_check = 0
resume = 0
rate_interval = "string"
no_ssl = 0
till = 0
flap_period = 0
sub_group = 0
dampening_flaps = 0
bw_rate_limit_no_logging = 0
down_grace_period = 0
initial_slow_start = 0
dscp = 0
request_rate_interval = "string"
add = 0
every = 0
shared_partition_pool = 0
conn_limit_no_logging = 0
extended_stats = 0
uuid = "string"
reset = 0
del_session_on_server_down = 0
conn_rate_limit_no_logging = 0
name = "string"
template_port_pool_shared = "string"
bw_rate_limit_duration = 0
bw_rate_limit_resume = 0
user_tag = "string"
times = 0
request_rate_no_logging = 0
down_timer = 0
conn_rate_limit = 0
source_nat = "string"
reassign = 0
health_check = "string"
 
}
```

## Argument Reference

* `add` - Slow start connection limit add by a number every interval (Add by this number every interval)
* `bw_rate_limit` - Configure bandwidth rate limit on real server port (Bandwidth rate limit in Kbps)
* `bw_rate_limit_duration` - Duration in seconds the observed rate needs to honor
* `bw_rate_limit_no_logging` - Do not log bandwidth rate limit related state transitions
* `bw_rate_limit_resume` - Resume server selection after bandwidth drops below this threshold (in Kbps) (Bandwidth rate limit resume threshold (in Kbps))
* `conn_limit` - Connection limit
* `conn_limit_no_logging` - Do not log connection over limit event
* `conn_rate_limit` - Connection rate limit
* `conn_rate_limit_no_logging` - Do not log connection over limit event
* `dampening_flaps` - service dampening flaps count (max-flaps allowed in flap period)
* `decrement` - Decrease after every round of DNS query (default is 0)
* `del_session_on_server_down` - Delete session if the server/port goes down (either disabled/hm down)
* `dest_nat` - Destination NAT
* `down_grace_period` - Port down grace period
* `down_timer` - The timer to bring the marked down server/port to up (default is 0, never bring up) (The timer to bring up server (in second, default is 0))
* `dscp` - Differentiated Services Code Point (DSCP to Real Server IP Mapping Value)
* `dynamic_member_priority` - Set dynamic member’s priority (Initial priority (default is 16))
* `every` - Slow start connection limit increment interval (default 10)
* `extended_stats` - Enable extended statistics on real server port
* `flap_period` - take service out of rotation if max-flaps exceeded within time in seconds
* `health_check` - Health Check Monitor (Health monitor name)
* `health_check_disable` - Disable configured health check configuration
* `inband_health_check` - Use inband traffic to detect port’s health status
* `initial_slow_start` - Initial slow start connection limit (default 128)
* `name` - Port template name
* `no_ssl` - No SSL
* `rate_interval` - ‘100ms’: Use 100 ms as sampling interval; ‘second’: Use 1 second as sampling interval;
* `reassign` - Maximum reassign times before declear the server/port down (default is 25) (The maximum reassign number)
* `request_rate_interval` - ‘100ms’: Use 100 ms as sampling interval; ‘second’: Use 1 second as sampling interval;
* `request_rate_limit` - Request rate limit
* `request_rate_no_logging` - Do not log connection over limit event
* `resel_on_reset` - When receiving reset from server, do the server/port reselection (default is 0, don’t do reselection)
* `reset` - Send client reset when connection rate over limit
* `restore_svc_time` - put the service back to the rotation after time in seconds
* `resume` - Resume accepting new connection after connection number drops below threshold (Connection resume threshold)
* `retry` - Maximum retry times before reassign this connection to another server/port (default is 2) (The maximum retry number)
* `shared_partition_pool` - Reference a NAT pool or pool-group from shared partition
* `slow_start` - Slowly ramp up the connection number after port is up
* `source_nat` - Source NAT (IP NAT Pool or pool group name)
* `stats_data_action` - ‘stats-data-enable’: Enable statistical data collection for real server port; ‘stats-data-disable’: Disable statistical data collection for real server port;
* `sub_group` - Divide service group members into different sub groups (Sub group ID (default is 0))
* `template_port_pool_shared` - Source NAT (IP NAT Pool or pool group name)
* `till` - Slow start ends when slow start connection limit reaches a number (default 4096) (Slow start ends when connection limit reaches this number)
* `times` - Slow start connection limit multiply by a number every interval (default 2) (Multiply by this number every interval)
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `weight` - Weight (port weight)
