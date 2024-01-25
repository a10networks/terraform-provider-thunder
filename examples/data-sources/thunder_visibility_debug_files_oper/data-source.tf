provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_visibility_debug_files_oper" "thunder_visibility_debug_files_oper" {

}
output "get_visibility_debug_files_oper" {
  value = ["${data.thunder_visibility_debug_files_oper.thunder_visibility_debug_files_oper}"]
}