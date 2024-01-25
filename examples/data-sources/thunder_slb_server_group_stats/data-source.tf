provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_server_group_stats" "thunder_slb_server_group_stats" {
  name = "server123"


}
output "get_slb_server_group_stats" {
  value = ["${data.thunder_slb_server_group_stats.thunder_slb_server_group_stats}"]
}