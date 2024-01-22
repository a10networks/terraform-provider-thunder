provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_service_group_stats" "thunder_slb_service_group_stats" {

  name = "sg180"
}
output "get_slb_service_group_stats" {
  value = ["${data.thunder_slb_service_group_stats.thunder_slb_service_group_stats}"]
}