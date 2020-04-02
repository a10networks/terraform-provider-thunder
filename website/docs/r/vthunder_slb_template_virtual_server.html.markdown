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
	conn_limit = 1
	conn_limit_reset = 1
	icmpv6_rate_limit = 1
	subnet_gratuitous_arp = 1
	tcp_stack_tfo_backoff_time = 1 
}
```

## Argument Reference

* `name` - Virtual server template name
* `user_tag` - Customized tag
* `conn_limit` - Connection limit
* `conn_limit_reset` - Send client reset when connection over limit
* `icmpv6_rate_limit` - ICMPv6 rate limit (Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit)
* `subnet_gratuitous_arp` - Send gratuitous ARP for every IP in the subnet virtual server
* `tcp_stack_tfo_backoff_time` - The time tcp stack will wait before allowing new fast-open requests after security condition, default 600 seconds (number)
