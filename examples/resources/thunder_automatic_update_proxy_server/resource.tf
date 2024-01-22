provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_automatic_update_proxy_server" "thunder_automatic_update_proxy_server" {
  auth_type     = "basic"
  https_port    = 16819
  password      = 1
  proxy_host    = "127"
  secret_string = "109"
  username      = "119"
}