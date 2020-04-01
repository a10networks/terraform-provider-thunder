---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_virtual_port"
sidebar_current: "docs-vthunder-resource-slb-template-virtual-port"
description: |-
    Provides details about vthunder slb template virtual-port resource for A10
---

# vthunder\_slb\_template\_virtual_port

`vthunder_slb_template_virtual_port provides details about slb template virtual-port
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_virtual_port" "testname" {
	name = "testvirtualport"
	user_tag = "test_tag"
	reset_unknown_conn = 0
	ignore_tcp_msl = 0
	rate = 5
	snat_msl = 10
	allow_syn_otherflags = 0
	aflow = 0
	conn_limit = 50
	drop_unknown_conn = 0
	reset_l7_on_failover = 0
	pkt_rate_type = "src-port"
	rate_interval = "100ms"
	snat_port_preserve = 0
	conn_rate_limit_reset = 0
	when_rr_enable = 0
	non_syn_initiation = 0
	conn_limit_reset = 0
	dscp = 50
	pkt_rate_limit_reset = 0
	conn_limit_no_logging = 0
	conn_rate_limit_no_logging = 0
	log_options = "no-logging"
	allow_vip_to_rport_mapping = 0
	pkt_rate_interval = "100ms"
	conn_rate_limit = 40
}
```

## Argument Reference

* `name` - Virtual port template name
* `user_tag` - Customized tag
* `reset_unknown_conn` - Send reset back if receives TCP packet without SYN or RST flag and it does not belong to any existing connections
* `ignore_tcp_msl` - reclaim TCP resource immediately without MSL
* `rate` - Source IP and port rate limit (Packet rate limit)
* `snat_msl` - Source NAT MSL (Source NAT MSL value (seconds))
* `allow_syn_otherflags` - Allow initial SYN packet with other flags
* `aflow` - Use aFlow to eliminate the traffic surge
* `conn_limit` - Connection limit
* `drop_unknown_conn` - Drop conection if receives TCP packet without SYN or RST flag and it does not belong to any existing connections
* `reset_l7_on_failover` - Send reset to L7 client and server connection upon a failover
* `pkt_rate_type` - 'src-ip-port’: Source IP and port rate limit; 'src-port’: Source port rate limit;
* `rate_interval` - '100ms’: Use 100 ms as sampling interval; 'second’: Use 1 second as sampling interval;
* `snat_port_preserve` - Source NAT Port Preservation
* `conn_rate_limit_reset` - Send client reset when connection rate over limit
* `when_rr_enable` - Only do rate limit if CPU RR triggered
* `non_syn_initiation` - Allow initial TCP packet to be non-SYN
* `conn_limit_reset` - Send client reset when connection over limit
* `dscp` - Differentiated Services Code Point (DSCP to Real Server IP Mapping Value)
* `pkt_rate_limit_reset` - send client-side reset (reset after packet limit)
* `conn_limit_no_logging` - Do not log connection over limit event
* `conn_rate_limit_no_logging` - Do not log connection over limit event
* `log_options` - 'no-logging’: Do not log over limit event; 'no-repeat-logging’: log once for over limit event. Default is log once per minute;
* `allow_vip_to_rport_mapping` - Allow mapping of VIP to real port
* `pkt_rate_interval` - '100ms’: Source IP and port rate limit per 100ms; 'second’: Source IP and port rate limit per second (default);
* `conn_rate_limit` - Connection rate limit

