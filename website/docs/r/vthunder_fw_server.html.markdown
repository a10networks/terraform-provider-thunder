---
layout: "vthunder"
page_title: "vthunder: vthunder_fw_server"
sidebar_current: "docs-vthunder-resource-fw-server"
description: |-
	Provides details about vthunder fw server resource for A10
---

# vthunder\_fw\_server

`vthunder_fw_server` Provides details about vthunder fw server
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_fw_server" "FwTest" {
	name = "a"
	server_ipv6_addr = "2003::1" 
}
```

## Argument Reference

* `action` - ‘enable’: enable; ‘disable’: disable;
* `fqdn_name` - Server hostname
* `health_check` - Health Check (Monitor Name)
* `health_check_disable` - Disable health check
* `host` - IP Address
* `name` - Server Name
* `resolve_as` - ‘resolve-to-ipv4’: Use A Query only to resolve FQDN; ‘resolve-to-ipv6’: Use AAAA Query only to resolve FQDN; ‘resolve-to-ipv4-and-ipv6’: Use A as well as AAAA Query to resolve FQDN;
* `server_ipv6_addr` - IPV6 address
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `port_number` - Port Number
* `protocol` - ‘tcp’: TCP Port; ‘udp’: UDP Port;
* `counters1` - ‘all’: all; ‘curr-conn’: Current connections; ‘total-conn’: Total connections; ‘fwd-pkt’: Forward packets; ‘rev-pkt’: Reverse Packets; ‘peak-conn’: Peak connections;

