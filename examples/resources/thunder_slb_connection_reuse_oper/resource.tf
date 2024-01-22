provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_connection_reuse_oper" "thunder_slb_connection_reuse_oper" {
}
output "get_slb_connection_reuse_oper" {
  value = ["${data.thunder_slb_connection_reuse_oper.thunder_slb_connection_reuse_oper}"]
}