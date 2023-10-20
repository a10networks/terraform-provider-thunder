provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_udp" "resourceSlbTemplateUdp" {
  name                       = "template_udp"
  qos                        = 1
  stateless_conn_timeout     = 5
  idle_timeout               = 1
  re_select_if_server_down   = 1
  disable_clear_session      = 1
  immediate                  = 1
  short                      = 0
  age                        = 0
  radius_lb_method_hash_type = "ip"
  avp                        = "4"
  user_tag                   = "udp_template"
}