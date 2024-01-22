provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dst_dynamic_entries_resource_usage_oper" "thunder_ddos_dst_dynamic_entries_resource_usage_oper" {
}
output "get_ddos_dst_dynamic_entries_resource_usage_oper" {
  value = ["${data.thunder_ddos_dst_dynamic_entries_resource_usage_oper.thunder_ddos_dst_dynamic_entries_resource_usage_oper}"]
}