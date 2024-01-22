provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_overlay_tunnel_vtep_stats" "thunder_overlay_tunnel_vtep_stats" {

  id1 = 56
}
output "get_overlay_tunnel_vtep_stats" {
  value = ["${data.thunder_overlay_tunnel_vtep_stats.thunder_overlay_tunnel_vtep_stats}"]
}