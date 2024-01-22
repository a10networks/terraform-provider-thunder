provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_server_stats" "thunder_slb_server_stats" {

  name = "test-server1"
}
output "get_slb_server_stats" {
  value = ["${data.thunder_slb_server_stats.thunder_slb_server_stats}"]
}