provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_reset_unknown_conn_oper" "thunder_slb_reset_unknown_conn_oper" {
}
output "get_slb_reset_unknown_conn_oper" {
  value = ["${data.thunder_slb_reset_unknown_conn_oper.thunder_slb_reset_unknown_conn_oper}"]
}