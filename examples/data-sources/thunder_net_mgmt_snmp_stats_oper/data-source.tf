provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_net_mgmt_snmp_stats_oper" "thunder_net_mgmt_snmp_stats_oper" {
}
output "get_net_mgmt_snmp_stats_oper" {
  value = ["${data.thunder_net_mgmt_snmp_stats_oper.thunder_net_mgmt_snmp_stats_oper}"]
}