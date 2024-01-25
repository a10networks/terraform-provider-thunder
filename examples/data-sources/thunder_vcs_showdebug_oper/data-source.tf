provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vcs_showdebug_oper" "thunder_vcs_showdebug_oper" {
}
output "get_vcs_showdebug_oper" {
  value = ["${data.thunder_vcs_showdebug_oper.thunder_vcs_showdebug_oper}"]
}