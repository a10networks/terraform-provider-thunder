provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vrrp_a_state_stats" "thunder_vrrp_a_state_stats" {
}
output "get_vrrp_a_state_stats" {
  value = ["${data.thunder_vrrp_a_state_stats.thunder_vrrp_a_state_stats}"]
}