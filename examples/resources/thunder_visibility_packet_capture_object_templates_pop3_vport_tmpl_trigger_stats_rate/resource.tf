provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_pop3_vport_tmpl_trigger_stats_rate" "thunder_visibility_packet_capture_object_templates_pop3_vport_tmpl_trigger_stats_rate" {

  name                  = "test"
  bad_sequence          = 1
  cl_est_err            = 1
  duration              = 60
  insert_tuple_fail     = 1
  invalid_start_line    = 1
  line_too_long         = 1
  no_route              = 1
  rsv_persist_conn_fail = 1
  ser_connecting_err    = 1
  server_response_err   = 1
  smp_v4_fail           = 1
  smp_v6_fail           = 1
  snat_fail             = 1
  svrsel_fail           = 1
  threshold_exceeded_by = 5
  unsupported_command   = 1
}