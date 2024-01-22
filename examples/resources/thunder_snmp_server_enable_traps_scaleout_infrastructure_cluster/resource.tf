provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_enable_traps_scaleout_infrastructure_cluster" "thunder_snmp_server_enable_traps_scaleout_infrastructure_cluster" {
  election                   = 0
  master_calling_re_election = 0
  node_status                = 0
}