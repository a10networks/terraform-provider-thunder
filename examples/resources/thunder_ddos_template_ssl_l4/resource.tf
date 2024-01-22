provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_ssl_l4" "thunder_ddos_template_ssl_l4" {

  action        = "drop"
  allow_non_tls = 1
  auth_config_cfg {
    timeout = 5
    trials  = 5
  }
  disable            = 1
  renegotiation      = 7
  request_rate_limit = 5738767
  ssl_l4_tmpl_name   = "test"
  ssl_traffic_check {
    header_inspection        = 1
    header_action            = "drop"
    check_resumed_connection = 1
  }
  user_tag = "87"

}