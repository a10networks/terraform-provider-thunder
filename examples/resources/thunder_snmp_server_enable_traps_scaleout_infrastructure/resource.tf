provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_enable_traps_scaleout_infrastructure" "thunder_snmp_server_enable_traps_scaleout_infrastructure" {
  all = 0
  cluster {
    election                   = 0
    master_calling_re_election = 0
    node_status                = 0
  }
  master_node {
    traffic_map_distribution   = 0
    vserver_traffic_map_update = 0
  }
  service_node {
    local_device_disabled = 0
    service_master        = 0
    traffic_map_update    = 0
  }
  test_send_all_traps = 0
}