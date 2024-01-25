provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_virtual_server_oper" "thunder_slb_virtual_server_oper" {

  name = "test"

}
output "get_slb_virtual_server_oper" {
  value = ["${data.thunder_slb_virtual_server_oper.thunder_slb_virtual_server_oper}"]
}