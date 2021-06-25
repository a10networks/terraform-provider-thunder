---
layout: "thunder"
page_title: "thunder: thunder_slb_template_virtual_port"
sidebar_current: "docs-thunder-resource-slb-template-virtual-port"
description: |-
	Provides details about thunder slb template virtual port resource for A10
---

# thunder\_slb\_template\_virtual\_port

`thunder_slb_template_virtual_port` Provides details about thunder slb template virtual port
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_slb_template_virtual_port" "Slb_Template_Virtual_Port_Test" {

reset_unknown_conn = 0
ignore_tcp_msl = 0
rate = 0
snat_msl = 0
allow_syn_otherflags = 0
aflow = 0
conn_limit = 0
drop_unknown_conn = 0
uuid = "string"
reset_l7_on_failover = 0
pkt_rate_type = "string"
rate_interval = "string"
snat_port_preserve = 0
conn_rate_limit_reset = 0
when_rr_enable = 0
non_syn_initiation = 0
conn_limit_reset = 0
dscp = 0
pkt_rate_limit_reset = 0
conn_limit_no_logging = 0
conn_rate_limit_no_logging = 0
log_options = "string"
name = "string"
allow_vip_to_rport_mapping = 0
pkt_rate_interval = "string"
user_tag = "string"
conn_rate_limit = 0
 
}
```

## Argument Reference

* `aflow` - Use aFlow to eliminate the traffic surge
* `allow_syn_otherflags` - Allow initial SYN packet with other flags
* `allow_vip_to_rport_mapping` - Allow mapping of VIP to real port
* `conn_limit` - Connection limit
* `conn_limit_no_logging` - Do not log connection over limit event
* `conn_limit_reset` - Send client reset when connection over limit
* `conn_rate_limit` - Connection rate limit
* `conn_rate_limit_no_logging` - Do not log connection over limit event
* `conn_rate_limit_reset` - Send client reset when connection rate over limit
* `drop_unknown_conn` - Drop conection if receives TCP packet without SYN or RST flag and it does not belong to any existing connections
* `dscp` - Differentiated Services Code Point (DSCP to Real Server IP Mapping Value)
* `ignore_tcp_msl` - reclaim TCP resource immediately without MSL
* `log_options` - ‘no-logging’: Do not log over limit event; ‘no-repeat-logging’: log once for over limit event. Default is log once per minute;
* `name` - Virtual port template name
* `non_syn_initiation` - Allow initial TCP packet to be non-SYN
* `pkt_rate_interval` - ‘100ms’: Source IP and port rate limit per 100ms; ‘second’: Source IP and port rate limit per second (default);
* `pkt_rate_limit_reset` - send client-side reset (reset after packet limit)
* `pkt_rate_type` - ‘src-ip-port’: Source IP and port rate limit; ‘src-port’: Source port rate limit;
* `rate` - Source IP and port rate limit (Packet rate limit)
* `rate_interval` - ‘100ms’: Use 100 ms as sampling interval; ‘second’: Use 1 second as sampling interval;
* `reset_l7_on_failover` - Send reset to L7 client and server connection upon a failover
* `reset_unknown_conn` - Send reset back if receives TCP packet without SYN or RST flag and it does not belong to any existing connections
* `snat_msl` - Source NAT MSL (Source NAT MSL value)
* `snat_port_preserve` - Source NAT Port Preservation
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `when_rr_enable` - Only do rate limit if CPU RR triggered
