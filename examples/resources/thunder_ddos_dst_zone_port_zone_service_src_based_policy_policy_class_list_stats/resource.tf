provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dst_zone_port_zone_service_src_based_policy_policy_class_list_stats" "thunder_ddos_dst_zone_port_zone_service_src_based_policy_policy_class_list_stats" {
  zone_name             = "test"
  port_num              = "20"
  protocol              = "tcp"
  src_based_policy_name = "test"
  class_list_name       = "test"

}
output "get_ddos_dst_zone_port_zone_service_src_based_policy_policy_class_list_stats" {
  value = ["${data.thunder_ddos_dst_zone_port_zone_service_src_based_policy_policy_class_list_stats.thunder_ddos_dst_zone_port_zone_service_src_based_policy_policy_class_list_stats}"]
}