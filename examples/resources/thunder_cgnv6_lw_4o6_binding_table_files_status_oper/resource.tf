provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_lw_4o6_binding_table_files_status_oper" "thunder_cgnv6_lw_4o6_binding_table_files_status_oper" {
}
output "get_cgnv6_lw_4o6_binding_table_files_status_oper" {
  value = ["${data.thunder_cgnv6_lw_4o6_binding_table_files_status_oper.thunder_cgnv6_lw_4o6_binding_table_files_status_oper}"]
}