provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_harmony_controller_tunnel_stats_oper" "thunder_harmony_controller_tunnel_stats_oper" {
}
output "get_harmony_controller_tunnel_stats_oper" {
  value = ["${data.thunder_harmony_controller_tunnel_stats_oper.thunder_harmony_controller_tunnel_stats_oper}"]
}