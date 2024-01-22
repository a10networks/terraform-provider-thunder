provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_tcp" "thunderSlbTemplateTcp" {
  name                       = "template_tcp"
  logging                    = "init"
  idle_timeout               = 1
  half_open_idle_timeout     = 1
  half_close_idle_timeout    = 60
  del_session_on_server_down = 1
  force_delete_timeout       = 1
  force_delete_timeout_100ms = 0
  initial_window_size        = 1
  insert_client_ip           = 1
  lan_fast_ack               = 1
  qos                        = 1
  re_select_if_server_down   = 1
  reset_follow_fin           = 1
  reset_fwd                  = 1
  reset_rev                  = 1
  disable                    = 1
  down                       = 0
  user_tag                   = "TCP_Template"
}