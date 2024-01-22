provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_ssl_l4_ssl_traffic_check" "thunder_ddos_zone_template_ssl_l4_ssl_traffic_check" {
  ssl_l4_tmpl_name         = "test"
  check_resumed_connection = 1
  header_action            = "drop"
  header_inspection        = 1
}