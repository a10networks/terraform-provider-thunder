

provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource  "thunder_snmp_server_enable_traps_routing_isis"   "SnMPServerEnableTrapsRoutingIsis"  {
    
    isis_adjacency_change= 1
    isis_area_mismatch= 1
    isis_attempt_to_exceed_max_sequence= 1
    isis_authentication_failure= 1
    isis_authentication_type_failure= 1
    isis_corrupted_lsp_detected= 1
    isis_database_overload= 1
    isis_id_len_mismatch= 1
    isis_lsp_too_large_to_propagate= 0
    isis_manual_address_drops= 1
    isis_max_area_addresses_mismatch= 1
    isis_originating_lsp_buffer_size_mismatch= 0
    isis_own_lsp_purge= 0
    isis_protocols_supported_mismatch= 0
    isis_rejected_adjacency= 0
    isis_sequence_number_skip= 0
    isis_version_skew= 0
   
}
