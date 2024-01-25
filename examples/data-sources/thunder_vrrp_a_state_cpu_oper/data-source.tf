provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vrrp_a_state_cpu_oper" "thunder_vrrp_a_state_cpu_oper" {
}
output "get_vrrp_a_state_cpu_oper" {
  value = ["${data.thunder_vrrp_a_state_cpu_oper.thunder_vrrp_a_state_cpu_oper}"]
}