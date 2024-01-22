provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_http_alg_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_http_alg_trigger_stats_inc" {

  name                     = "test"
  header_insertion_failed  = 1
  header_removal_failed    = 1
  out_of_memory_dropped    = 1
  out_of_order_dropped     = 1
  queue_len_exceed_dropped = 1
  radius_requst_dropped    = 1
  radius_response_dropped  = 1
}