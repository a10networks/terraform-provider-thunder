provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_tcp_syn_per_sec_oper" "thunder_system_tcp_syn_per_sec_oper" {
}
output "get_system_tcp_syn_per_sec_oper" {
  value = ["${data.thunder_system_tcp_syn_per_sec_oper.thunder_system_tcp_syn_per_sec_oper}"]
}