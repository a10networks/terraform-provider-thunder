provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_web_service" "thunder_web_service" {
  auto_redirt_disable     = 1
  axapi_idle              = 1
  axapi_session_limit     = 100
  gui_idle                = 1
  gui_session_limit       = 100
  port                    = 82
  login_message           = "welcome"
  pre_login_message       = "NowWelcome"
  max_keep_alive_requests = 101
  keep_alive_timeout      = 40
  mpm_max_conn            = 2
  mpm_min_spare_conn      = 1
  mpm_max_conn_per_child  = 22

  secure_port           = 443
  secure_server_disable = 1
  server_disable        = 1
}