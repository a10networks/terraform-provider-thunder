---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_enable_traps_routing_ospf"
sidebar_current: "docs-thunder-resource-snmp-server-enable-traps-routing-ospf"
description: |-
	Provides details about thunder snmp server enable traps routing ospf resource for A10
---

# thunder\_snmp\_server\_enable\_traps\_routing\_ospf

`thunder_snmp_server_enable_traps_routing_ospf` Provides details about thunder snmp server enable traps routing ospf
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_snmp_server_enable_traps_routing_ospf" "Snmp_Server_Enable_Traps_Routing_Ospf_Test" {

ospf_lsdb_overflow = 0
uuid = "string"
ospf_nbr_state_change = 0
ospf_if_state_change = 0
ospf_virt_nbr_state_change = 0
ospf_lsdb_approaching_overflow = 0
ospf_if_auth_failure = 0
ospf_virt_if_auth_failure = 0
ospf_virt_if_config_error = 0
ospf_virt_if_rx_bad_packet = 0
ospf_tx_retransmit = 0
ospf_virt_if_state_change = 0
ospf_if_config_error = 0
ospf_max_age_lsa = 0
ospf_if_rx_bad_packet = 0
ospf_virt_if_tx_retransmit = 0
ospf_originate_lsa = 0
 
}
```

## Argument Reference

* `ospf_if_auth_failure` - Enable ospfIfAuthFailure traps
* `ospf_if_config_error` - Enable ospfIfConfigError traps
* `ospf_if_rx_bad_packet` - Enable ospfIfRxBadPacket traps
* `ospf_if_state_change` - Enable ospfIfStateChange traps
* `ospf_lsdb_approaching_overflow` - Enable ospfLsdbApproachingOverflow traps
* `ospf_lsdb_overflow` - Enable ospfLsdbOverflow traps
* `ospf_max_age_lsa` - Enable ospfMaxAgeLsa traps
* `ospf_nbr_state_change` - Enable ospfNbrStateChange traps
* `ospf_originate_lsa` - Enable ospfOriginateLsa traps
* `ospf_tx_retransmit` - Enable ospfTxRetransmit traps
* `ospf_virt_if_auth_failure` - Enable ospfVirtIfAuthFailure traps
* `ospf_virt_if_config_error` - Enable ospfVirtIfConfigError traps
* `ospf_virt_if_rx_bad_packet` - Enable ospfVirtIfRxBadPacket traps
* `ospf_virt_if_state_change` - Enable ospfVirtIfStateChange traps
* `ospf_virt_if_tx_retransmit` - Enable ospfVirtIfTxRetransmit traps
* `ospf_virt_nbr_state_change` - Enable ospfVirtNbrStateChange traps
* `uuid` - uuid of the object

