provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_visibility_resource_usage_oper" "thunder_visibility_resource_usage_oper" {

}
output "get_visibility_resource_usage_oper" {
  value = ["${data.thunder_visibility_resource_usage_oper.thunder_visibility_resource_usage_oper}"]
}