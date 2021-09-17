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
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_slb_template_virtual_port" "resourceSlbTemplateVirtualPortTest" {
	name = "string"
aflow = 0
allow_syn_otherflags = 0
conn_limit = 0
conn_limit_reset = 0
conn_limit_no_logging = 0
conn_rate_limit = 0
rate_interval = "string"
conn_rate_limit_reset = 0
conn_rate_limit_no_logging = 0
pkt_rate_type = "string"
rate = 0
pkt_rate_interval = "string"
pkt_rate_limit_reset = 0
log_options = "string"
when_rr_enable = 0
allow_vip_to_rport_mapping = 0
dscp = 0
drop_unknown_conn = 0
reset_unknown_conn = 0
reset_l7_on_failover = 0
ignore_tcp_msl = 0
snat_msl = 0
snat_port_preserve = 0
non_syn_initiation = 0
uuid = "string"
user_tag = "string"
 
}

```

## Argument Reference

* `virtual-port` - Virtual port template
* `name` - Virtual port template name
* `aflow` - Use aFlow to eliminate the traffic surge
* `allow-syn-otherflags` - Allow initial SYN packet with other flags
* `conn-limit` - Connection limit
* `conn-limit-reset` - Send client reset when connection over limit
* `conn-limit-no-logging` - Do not log connection over limit event
* `conn-rate-limit` - Connection rate limit
* `rate-interval` - '100ms': Use 100 ms as sampling interval; 'second': Use 1 second as sampling interval;
* `conn-rate-limit-reset` - Send client reset when connection rate over limit
* `conn-rate-limit-no-logging` - Do not log connection over limit event
* `pkt-rate-type` - 'src-ip-port': Source IP and port rate limit; 'src-port': Source port rate limit;
* `rate` - Source IP and port rate limit (Packet rate limit)
* `pkt-rate-interval` - '100ms': Source IP and port rate limit per 100ms; 'second': Source IP and port rate limit per second (default);
* `pkt-rate-limit-reset` - send client-side reset (reset after packet limit)
* `log-options` - 'no-logging': Do not log over limit event; 'no-repeat-logging': log once for over limit event. Default is log once per minute;
* `when-rr-enable` - Only do rate limit if CPU RR triggered
* `allow-vip-to-rport-mapping` - Allow mapping of VIP to real port
* `dscp` - Differentiated Services Code Point (DSCP to Real Server IP Mapping Value)
* `drop-unknown-conn` - Drop conection if receives TCP packet without SYN or RST flag and it does not belong to any existing connections
* `reset-unknown-conn` - Send reset back if receives TCP packet without SYN or RST flag and it does not belong to any existing connections
* `reset-l7-on-failover` - Send reset to L7 client and server connection upon a failover
* `ignore-tcp-msl` - reclaim TCP resource immediately without MSL
* `snat-msl` - Source NAT MSL (Source NAT MSL value (seconds))
* `snat-port-preserve` - Source NAT Port Preservation
* `non-syn-initiation` - Allow initial TCP packet to be non-SYN
* `uuid` - uuid of the object
* `user-tag` - Customized tag

