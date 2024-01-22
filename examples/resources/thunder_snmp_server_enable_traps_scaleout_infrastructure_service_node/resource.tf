provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_enable_traps_scaleout_infrastructure_service_node" "thunder_snmp_server_enable_traps_scaleout_infrastructure_service_node" {
  local_device_disabled = 0
  service_master        = 0
  traffic_map_update    = 0
}