provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_scaleout_debug_redirect_table_l3_oper" "thunder_scaleout_debug_redirect_table_l3_oper" {
}
output "get_scaleout_debug_redirect_table_l3_oper" {
  value = ["${data.thunder_scaleout_debug_redirect_table_l3_oper.thunder_scaleout_debug_redirect_table_l3_oper}"]
}