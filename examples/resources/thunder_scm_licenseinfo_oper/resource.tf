provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_scm_licenseinfo_oper" "thunder_scm_licenseinfo_oper" {
}
output "get_scm_licenseinfo_oper" {
  value = ["${data.thunder_scm_licenseinfo_oper.thunder_scm_licenseinfo_oper}"]
}