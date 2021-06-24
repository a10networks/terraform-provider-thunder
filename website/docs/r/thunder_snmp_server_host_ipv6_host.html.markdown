---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_host_ipv6_host"
sidebar_current: "docs-thunder-resource-snmp-server-host-ipv6-host"
description: |-
	Provides details about thunder snmp server host ipv6 host resource for A10
---

# thunder\_snmp\_server\_host\_ipv6\_host

`thunder_snmp_server_host_ipv6_host` Provides details about thunder snmp server host ipv6 host
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}



resource "thunder_snmp_server_host_ipv6_host" "resourceSnmpServerHostIpv6HostTest" {
	uuid = "string"
ipv6_addr = "string"
udp_port = 0
v1_v2c_comm = "string"
user = "string"
version = "string"
 
}

```

## Argument Reference

* `ipv6_addr` - help IPV6 address of SNMP trap host
* `udp_port` - The trap host’s UDP port number(default: 162)
* `user` - SNMPv3 user to send traps (User Name)
* `uuid` - uuid of the object
* `v1_v2c_comm` - SNMPv1/v2c community string
* `version` - ‘v1’: Use SNMPv1; ‘v2c’: Use SNMPv2c; ‘v3’: User SNMPv3;

