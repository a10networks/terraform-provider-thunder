provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dst_entry_src_dst_pair_policy_policy_class_list_stats" "thunder_ddos_dst_entry_src_dst_pair_policy_policy_class_list_stats" {
  class_list_name       = "test"
  dst_entry_name        = "test"
  src_based_policy_name = "test"


}

output "get_ddos_dst_entry_src_dst_pair_policy_policy_class_list_stats" {
  value = ["${data.thunder_ddos_dst_entry_src_dst_pair_policy_policy_class_list_stats.thunder_ddos_dst_entry_src_dst_pair_policy_policy_class_list_stats}"]
}