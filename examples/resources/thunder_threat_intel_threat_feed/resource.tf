provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_threat_intel_threat_feed" "thunder_threat_intel_threat_feed" {
  type               = "webroot"
  log_level          = "warning"
  port               = 443
  proxy_auth_type    = "ntlm"
  proxy_host         = "44"
  proxy_password     = 1
  proxy_port         = 23599
  proxy_username     = "82"
  rtu_update_disable = 1
  secret_string      = "87"
  server_timeout     = 15
  update_interval    = 120
  user_tag           = "47"
}