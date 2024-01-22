provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_auth" "thunder_debug_auth" {
  authd          = 1
  client_addr    = "10.10.10.10"
  level          = "1"
  saml           = 1
  saml_sp        = "43"
  username       = "test"
  virtual_server = "vserver"
}