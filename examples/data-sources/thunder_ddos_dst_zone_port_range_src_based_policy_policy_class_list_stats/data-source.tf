provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dst_zone_port_range_src_based_policy_policy_class_list_stats" "thunder_ddos_dst_zone_port_range_src_based_policy_policy_class_list_stats" {
  class_list_name       = "test"
  zone_name             = "test"
  port_range_start      = 20
  port_range_end        = 22
  protocol              = "tcp"
  src_based_policy_name = "test"

}
output "get_ddos_dst_zone_port_range_src_based_policy_policy_class_list_stats" {
  value = ["${data.thunder_ddos_dst_zone_port_range_src_based_policy_policy_class_list_stats.thunder_ddos_dst_zone_port_range_src_based_policy_policy_class_list_stats}"]
}