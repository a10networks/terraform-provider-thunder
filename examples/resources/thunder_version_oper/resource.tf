provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_version_oper" "thunder_version_oper" {
}
output "get_version_oper" {
  value = ["${data.thunder_version_oper.thunder_version_oper}"]
}