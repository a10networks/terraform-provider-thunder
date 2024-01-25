provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_scm_license_src_info_oper" "thunder_scm_license_src_info_oper" {
}
output "get_scm_license_src_info_oper" {
  value = ["${data.thunder_scm_license_src_info_oper.thunder_scm_license_src_info_oper}"]
}