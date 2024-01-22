provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_enable_traps_lsn" "thunder_snmp_server_enable_traps_lsn" {
  all                                = 0
  fixed_nat_port_mapping_file_change = 1
  max_ipport_threshold               = 64512
  max_port_threshold                 = 655350000
  per_ip_port_usage_threshold        = 1
  total_port_usage_threshold         = 1
  traffic_exceeded                   = 1
}