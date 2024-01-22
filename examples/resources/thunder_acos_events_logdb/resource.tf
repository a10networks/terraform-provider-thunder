provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_acos_events_logdb" "thunder_acos_events_logdb" {
  enable_all                = 0
  enable_cgn                = 0
  enable_fw                 = 0
  enable_http_forward_proxy = 0
  enable_link_cost          = 0
  enable_mqtt               = 0
  enable_smtp               = 0
  enable_ssli               = 0
}