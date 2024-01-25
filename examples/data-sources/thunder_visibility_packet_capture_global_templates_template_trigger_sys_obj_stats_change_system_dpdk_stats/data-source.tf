provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_dpdk_stats" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_dpdk_stats" {

  name = "test"
}
output "get_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_dpdk_stats" {
  value = ["${data.thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_dpdk_stats.thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_dpdk_stats}"]
}