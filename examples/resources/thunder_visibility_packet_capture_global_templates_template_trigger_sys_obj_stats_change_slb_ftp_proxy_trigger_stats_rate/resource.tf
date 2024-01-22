provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_ftp_proxy_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_ftp_proxy_trigger_stats_rate" {

  name                     = "test"
  auth_fail                = 1
  bad_sequence             = 1
  cant_find_eprt           = 1
  cant_find_port           = 1
  cl_est_err               = 1
  cl_request_err           = 1
  data_conn_start_err      = 1
  data_send_fail           = 1
  data_serv_connected_err  = 1
  data_serv_connecting_err = 1
  data_server_conn_fail    = 1
  ds_fail                  = 1
  duration                 = 60
  insert_tuple_fail        = 1
  invalid_start_line       = 1
  line_too_long            = 1
  no_route                 = 1
  rsv_persist_conn_fail    = 1
  ser_connecting_err       = 1
  server_response_err      = 1
  smp_create_fail          = 1
  smp_v4_fail              = 1
  smp_v6_fail              = 1
  snat_fail                = 1
  svrsel_fail              = 1
  threshold_exceeded_by    = 5
  unsupported_command      = 1
  unsupported_pbsz_value   = 1
  unsupported_prot_value   = 1
}