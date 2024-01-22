provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_ip_threat_action" "thunder_template_ip_threat_action" {
  idle_timeout    = 5
  template_number = 2
  user_tag        = "124"
}