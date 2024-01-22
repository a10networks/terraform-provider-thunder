provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_hardware_accelerate_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_hardware_accelerate_trigger_stats_rate" {

  name                         = "test"
  duration                     = 60
  hw_fwd_flow_error_count      = 1
  hw_fwd_flow_seq_mismatch     = 1
  hw_fwd_flow_singlebit_errors = 1
  hw_fwd_flow_tag_mismatch     = 1
  hw_fwd_flow_unalign_count    = 1
  hw_fwd_flow_underflow_count  = 1
  hw_fwd_prog_errors           = 1
  threshold_exceeded_by        = 5
}