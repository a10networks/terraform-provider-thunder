provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_interface_trunk_stats" "thunder_interface_trunk_stats" {

  ifnum = 11

}
output "get_interface_trunk_stats" {
  value = ["${data.thunder_interface_trunk_stats.thunder_interface_trunk_stats}"]
}