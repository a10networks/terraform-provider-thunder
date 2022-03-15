---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_snmp_server_host_ipv6_host Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_snmp_server_host_ipv6_host: Specify hosts to receive SNMP traps
  _
---

# thunder_snmp_server_host_ipv6_host (Resource)

`thunder_snmp_server_host_ipv6_host`: Specify hosts to receive SNMP traps

_



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **ipv6_addr** (String) IPv6 address of SNMP trap host
- **version** (String) 'v1': Use SNMPv1; 'v2c': Use SNMPv2c; 'v3': User SNMPv3;

### Optional

- **id** (String) The ID of this resource.
- **udp_port** (Number) The trap host's UDP port number(default: 162)
- **user** (String) SNMPv3 user to send traps (User Name)
- **uuid** (String) uuid of the object
- **v1_v2c_comm** (String) SNMPv1/v2c community string

