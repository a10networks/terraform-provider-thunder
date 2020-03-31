---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_virtual_server"
sidebar_current: "docs-vthunder-resource-slb-template-virtual-server"
description: |-
    Provides details about vthunder slb template virtual-server resource for A10
---

# vthunder\_slb\_template\_virtual_server

`vthunder_slb_template_virtual_server provides details about slb template virtual-server
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_virtual_server" "testname" {
	name = "testvirtualserver"
	user_tag = "test_tag"
	conn_limit = 10000
	conn_rate_limit_no_logging = 0
	icmp_lockup_period = 10
	conn_limit_reset = 0
	rate_interval = "100ms"
	icmpv6_rate_limit = 50
	subnet_gratuitous_arp = 0
	icmpv6_lockup = 2000
	conn_rate_limit_reset = 1
	tcp_stack_tfo_backoff_time = 60
	tcp_stack_tfo_cookie_time_limit = 60
	conn_limit_no_logging = 0
	icmpv6_lockup_period = 50
	conn_rate_limit = 50
	tcp_stack_tfo_active_conn_limit = 50
	icmp_lockup = 50
	icmp_rate_limit = 50
}
```

## Argument Reference

* `name` - Virtual server template name
* `user_tag` - Customized tag
* `conn_limit` - Connection limit
* `conn_rate_limit_no_logging` - Do not log connection over limit event
* `icmp_lockup_period` - Lockup period (second)
* `conn_limit_reset` - Send client reset when connection over limit
* `rate_interval` - '100ms’: Use 100 ms as sampling interval; 'second’: Use 1 second as sampling interval;
* `icmpv6_rate_limit` - ICMPv6 rate limit (Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit)
* `subnet_gratuitous_arp` - Send gratuitous ARP for every IP in the subnet virtual server
* `icmpv6_lockup` - Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)
* `conn_rate_limit_reset` - Send client reset when connection rate over limit
* `tcp_stack_tfo_backoff_time` - The time tcp stack will wait before allowing new fast-open requests after security condition, default 600 seconds (number)
* `tcp_stack_tfo_cookie_time_limit` - The time limit (in seconds) that a layer 7 tcp fast-open cookie is valid, default is 60 seconds (number)
* `conn_limit_no_logging` - Do not log connection over limit event
* `icmpv6_lockup_period` - Lockup period (second)
* `conn_rate_limit` - Connection rate limit
* `tcp_stack_tfo_active_conn_limit` - The allowed active layer 7 tcp fast-open connection limit, default is zero (number)
* `icmp_lockup` - Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)
* `icmp_rate_limit` - ICMP rate limit (Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit)
