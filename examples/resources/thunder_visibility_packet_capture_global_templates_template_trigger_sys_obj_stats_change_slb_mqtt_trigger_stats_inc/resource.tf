provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_mqtt_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_mqtt_trigger_stats_inc" {

  name                   = "test"
  client_id_null         = 1
  conn_null              = 1
  insertion_failed       = 1
  parse_connect_fail     = 1
  parse_publish_fail     = 1
  parse_subscribe_fail   = 1
  parse_unsubscribe_fail = 1
  session_exist          = 1
  tuple_already_linked   = 1
  tuple_not_linked       = 1
}