---
layout: "thunder"
page_title: "thunder: thunder_slb_template_virtual_server"
sidebar_current: "docs-thunder-resource-slb-template-virtual-server"
description: |-
	Provides details about thunder slb template virtual server resource for A10
---

# thunder\_slb\_template\_virtual\_server

`thunder_slb_template_virtual_server` Provides details about thunder slb template virtual server
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_slb_template_virtual_server" "Slb_Template_Virtual_Server_Test" {

tcp_stack_tfo_cookie_time_limit = 0
subnet_gratuitous_arp = 0
icmpv6_lockup = 0
icmp_lockup_period = 0
icmpv6_lockup_period = 0
disable_when_all_ports_down = 0
conn_limit = 0
disable_when_any_port_down = 0
uuid = "string"
rate_interval = "string"
tcp_stack_tfo_active_conn_limit = 0
conn_rate_limit_reset = 0
icmp_lockup = 0
conn_limit_reset = 0
conn_limit_no_logging = 0
conn_rate_limit_no_logging = 0
name = "string"
icmpv6_rate_limit = 0
user_tag = "string"
conn_rate_limit = 0
tcp_stack_tfo_backoff_time = 0
icmp_rate_limit = 0
 
}
```

## Argument Reference

* `conn_limit` - Connection limit
* `conn_limit_no_logging` - Do not log connection over limit event
* `conn_limit_reset` - Send client reset when connection over limit
* `conn_rate_limit` - Connection rate limit
* `conn_rate_limit_no_logging` - Do not log connection over limit event
* `conn_rate_limit_reset` - Send client reset when connection rate over limit
* `icmp_lockup` - Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)
* `icmp_lockup_period` - Lockup period (second)
* `icmp_rate_limit` - ICMP rate limit (Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit)
* `icmpv6_lockup` - Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)
* `icmpv6_lockup_period` - Lockup period (second)
* `icmpv6_rate_limit` - ICMPv6 rate limit (Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit)
* `name` - Virtual server template name
* `rate_interval` - ‘100ms’: Use 100 ms as sampling interval; ‘second’: Use 1 second as sampling interval;
* `subnet_gratuitous_arp` - Send gratuitous ARP for every IP in the subnet virtual server
* `tcp_stack_tfo_active_conn_limit` - The allowed active layer 7 tcp fast-open connection limit, default is zero (number)
* `tcp_stack_tfo_backoff_time` - The time tcp stack will wait before allowing new fast-open requests after security condition, default 600 seconds (number)
* `tcp_stack_tfo_cookie_time_limit` - The time limit (in seconds) that a layer 7 tcp fast-open cookie is valid, default is 60 seconds (number)
* `user_tag` - Customized tag
* `uuid` - uuid of the object
