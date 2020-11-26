---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_host_ipv4_host"
sidebar_current: "docs-thunder-resource-snmp-server-host-ipv4-host"
description: |-
	Provides details about thunder snmp server host ipv4 host resource for A10
---

# thunder\_snmp\_server\_host\_ipv4\_host

`thunder_snmp_server_host_ipv4_host` Provides details about thunder snmp server host ipv4 host
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_snmp_server_host_ipv4_host" "resourceSnmpServerHostIpv4HostTest" {
	uuid = "string"
ipv4_addr = "string"
udp_port = 0
v1_v2c_comm = "string"
user = "string"
version = "string"
 
}

```

## Argument Reference

* `ipv4_addr` - help IPV4 address of SNMP trap host
* `udp_port` - The trap host’s UDP port number(default: 162)
* `user` - SNMPv3 user to send traps (User Name)
* `uuid` - uuid of the object
* `v1_v2c_comm` - SNMPv1/v2c community string
* `version` - ‘v1’: Use SNMPv1; ‘v2c’: Use SNMPv2c; ‘v3’: User SNMPv3;

