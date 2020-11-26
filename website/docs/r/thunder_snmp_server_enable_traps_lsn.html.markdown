---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_enable_traps_lsn"
sidebar_current: "docs-thunder-resource-snmp-server-enable-traps-lsn"
description: |-
	Provides details about thunder snmp server enable traps lsn resource for A10
---

# thunder\_snmp\_server\_enable\_traps\_lsn

`thunder_snmp_server_enable_traps_lsn` Provides details about thunder snmp server enable traps lsn
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_snmp_server_enable_traps_lsn" "Snmp_Server_Enable_Traps_Lsn_Test" {
all = 0
fixed_nat_port_mapping_file_change = 0
per_ip_port_usage_threshold = 0
uuid = "string"
total_port_usage_threshold = 0
max_port_threshold = 0
max_ipport_threshold = 0
traffic_exceeded = 0
 
}


```

## Argument Reference

* `all` - Enable all LSN group traps
* `fixed_nat_port_mapping_file_change` - Enable LSN trap when fixed nat port mapping file change
* `max_ipport_threshold` - Maximum threshold
* `max_port_threshold` - Maximum threshold
* `per_ip_port_usage_threshold` - Enable LSN trap when IP total port usage reaches the threshold (default 64512)
* `total_port_usage_threshold` - Enable LSN trap when NAT total port usage reaches the threshold (default 655350000)
* `traffic_exceeded` - Enable LSN trap when NAT pool reaches the threshold
* `uuid` - uuid of the object

