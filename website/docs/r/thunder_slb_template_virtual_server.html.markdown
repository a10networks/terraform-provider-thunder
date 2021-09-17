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
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_slb_template_virtual_server" "resourceSlbTemplateVirtualServerTest" {
	name = "string"
conn_limit = 0
conn_limit_reset = 0
conn_limit_no_logging = 0
conn_rate_limit = 0
rate_interval = "string"
conn_rate_limit_reset = 0
conn_rate_limit_no_logging = 0
icmp_rate_limit = 0
icmp_lockup = 0
icmp_lockup_period = 0
icmpv6_rate_limit = 0
icmpv6_lockup = 0
icmpv6_lockup_period = 0
tcp_stack_tfo_active_conn_limit = 0
tcp_stack_tfo_cookie_time_limit = 0
tcp_stack_tfo_backoff_time = 0
subnet_gratuitous_arp = 0
disable_when_all_ports_down = 0
disable_when_any_port_down = 0
uuid = "string"
user_tag = "string"
 
}

```

## Argument Reference

* `virtual-server` - Virtual server template
* `name` - Virtual server template name
* `conn-limit` - Connection limit
* `conn-limit-reset` - Send client reset when connection over limit
* `conn-limit-no-logging` - Do not log connection over limit event
* `conn-rate-limit` - Connection rate limit
* `rate-interval` - '100ms': Use 100 ms as sampling interval; 'second': Use 1 second as sampling interval;
* `conn-rate-limit-reset` - Send client reset when connection rate over limit
* `conn-rate-limit-no-logging` - Do not log connection over limit event
* `icmp-rate-limit` - ICMP rate limit (Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit)
* `icmp-lockup` - Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)
* `icmp-lockup-period` - Lockup period (second)
* `icmpv6-rate-limit` - ICMPv6 rate limit (Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit)
* `icmpv6-lockup` - Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)
* `icmpv6-lockup-period` - Lockup period (second)
* `tcp-stack-tfo-active-conn-limit` - The allowed active layer 7 tcp fast-open connection limit, default is zero (number)
* `tcp-stack-tfo-cookie-time-limit` - The time limit (in seconds) that a layer 7 tcp fast-open cookie is valid, default is 60 seconds (number)
* `tcp-stack-tfo-backoff-time` - The time tcp stack will wait before allowing new fast-open requests after security condition, default 600 seconds (number)
* `subnet-gratuitous-arp` - Send gratuitous ARP for every IP in the subnet virtual server
* `disable-when-all-ports-down` - Disable Virtual Server when all member ports are down
* `disable-when-any-port-down` - Disable Virtual Server when any member port is down
* `uuid` - uuid of the object
* `user-tag` - Customized tag

