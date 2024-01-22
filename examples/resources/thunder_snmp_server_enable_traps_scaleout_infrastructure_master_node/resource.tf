provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_enable_traps_scaleout_infrastructure_master_node" "thunder_snmp_server_enable_traps_scaleout_infrastructure_master_node" {
  traffic_map_distribution   = 0
  vserver_traffic_map_update = 0
}