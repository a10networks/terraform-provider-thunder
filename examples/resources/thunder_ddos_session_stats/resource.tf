provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_session_stats" "thunder_ddos_session_stats" {
}
output "get_ddos_session_stats" {
  value = ["${data.thunder_ddos_session_stats.thunder_ddos_session_stats}"]
}