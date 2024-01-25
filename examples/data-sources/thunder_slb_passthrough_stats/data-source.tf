provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_passthrough_stats" "thunder_slb_passthrough_stats" {
}
output "get_slb_passthrough_stats" {
  value = ["${data.thunder_slb_passthrough_stats.thunder_slb_passthrough_stats}"]
}