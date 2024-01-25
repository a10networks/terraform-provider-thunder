provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_scm_licensecpe_oper" "thunder_scm_licensecpe_oper" {
}
output "get_scm_licensecpe_oper" {
  value = ["${data.thunder_scm_licensecpe_oper.thunder_scm_licensecpe_oper}"]
}