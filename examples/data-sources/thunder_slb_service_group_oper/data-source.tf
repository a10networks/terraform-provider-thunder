provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_service_group_oper" "thunder_slb_service_group_oper" {

  name = "sg180"
}
output "get_slb_service_group_oper" {
  value = ["${data.thunder_slb_service_group_oper.thunder_slb_service_group_oper}"]
}