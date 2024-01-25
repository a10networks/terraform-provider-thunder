provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_ip_threat_list_stats" "thunder_system_ip_threat_list_stats" {
}
output "get_system_ip_threat_list_stats" {
  value = ["${data.thunder_system_ip_threat_list_stats.thunder_system_ip_threat_list_stats}"]
}