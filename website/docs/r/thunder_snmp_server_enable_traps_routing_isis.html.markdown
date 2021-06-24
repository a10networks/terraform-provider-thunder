---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_enable_traps_routing_isis"
sidebar_current: "docs-thunder-resource-snmp-server-enable-traps-routing-isis"
description: |-
	Provides details about thunder snmp server enable traps routing isis resource for A10
---

# thunder\_snmp\_server\_enable\_traps\_routing\_isis

`thunder_snmp_server_enable_traps_routing_isis` Provides details about thunder snmp server enable traps routing isis
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_snmp_server_enable_traps_routing_isis" "Snmp_Server_Enable_Traps_Routing_Isis_Test" {

isis_authentication_failure = 0
uuid = "string"
isis_protocols_supported_mismatch = 0
isis_rejected_adjacency = 0
isis_max_area_addresses_mismatch = 0
isis_corrupted_lsp_detected = 0
isis_originating_lsp_buffer_size_mismatch = 0
isis_area_mismatch = 0
isis_lsp_too_large_to_propagate = 0
isis_own_lsp_purge = 0
isis_sequence_number_skip = 0
isis_database_overload = 0
isis_attempt_to_exceed_max_sequence = 0
isis_id_len_mismatch = 0
isis_authentication_type_failure = 0
isis_version_skew = 0
isis_manual_address_drops = 0
isis_adjacency_change = 0
 
}
```

## Argument Reference

* `isis_adjacency_change` - Enable isisAdjacencyChange traps
* `isis_area_mismatch` - Enable isisAreaMismatch traps
* `isis_attempt_to_exceed_max_sequence` - Enable isisAttemptToExceedMaxSequence traps
* `isis_authentication_failure` - Enable isisAuthenticationFailure traps
* `isis_authentication_type_failure` - Enable isisAuthenticationTypeFailure traps
* `isis_corrupted_lsp_detected` - Enable isisCorruptedLSPDetected traps
* `isis_database_overload` - Enable isisDatabaseOverload traps
* `isis_id_len_mismatch` - Enable isisIDLenMismatch traps
* `isis_lsp_too_large_to_propagate` - Enable isisLSPTooLargeToPropagate traps
* `isis_manual_address_drops` - Enable isisManualAddressDrops traps
* `isis_max_area_addresses_mismatch` - Enable isisMaxAreaAddressesMismatch traps
* `isis_originating_lsp_buffer_size_mismatch` - Enable isisOriginatingLSPBufferSizeMismatch traps
* `isis_own_lsp_purge` - Enable isisOwnLSPPurge traps
* `isis_protocols_supported_mismatch` - Enable isisProtocolsSupportedMismatch traps
* `isis_rejected_adjacency` - Enable isisRejectedAdjacency traps
* `isis_sequence_number_skip` - Enable isisSequenceNumberSkip traps
* `isis_version_skew` - Enable isisVersionSkew traps
* `uuid` - uuid of the object

