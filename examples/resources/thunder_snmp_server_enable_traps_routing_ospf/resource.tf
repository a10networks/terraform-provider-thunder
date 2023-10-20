
provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource  "thunder_snmp_server_enable_traps_routing_ospf"   "SnMPServerEnableTrapsRoutingOspf"  {

    ospf_if_auth_failure= 1
    ospf_if_config_error= 0
    ospf_if_rx_bad_packet= 1
    ospf_if_state_change= 0
    ospf_lsdb_approaching_overflow= 1
    ospf_lsdb_overflow= 1
    ospf_max_age_lsa= 0
    ospf_nbr_state_change= 1
    ospf_originate_lsa= 0
    ospf_tx_retransmit= 0
    ospf_virt_if_auth_failure= 1
    ospf_virt_if_config_error= 1
    ospf_virt_if_rx_bad_packet= 0
    ospf_virt_if_state_change= 0
    ospf_virt_if_tx_retransmit= 0
    ospf_virt_nbr_state_change= 0
      
  }
  