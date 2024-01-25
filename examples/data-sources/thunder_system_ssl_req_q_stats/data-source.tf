provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_ssl_req_q_stats" "thunder_system_ssl_req_q_stats" {
}
output "get_system_ssl_req_q_stats" {
  value = ["${data.thunder_system_ssl_req_q_stats.thunder_system_ssl_req_q_stats}"]
}