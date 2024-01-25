provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_aam_resource_usage_oper" "thunder_aam_resource_usage_oper" {
}
output "get_aam_resource_usage_oper" {
  value = ["${data.thunder_aam_resource_usage_oper.thunder_aam_resource_usage_oper}"]
}