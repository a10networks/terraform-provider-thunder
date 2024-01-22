provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_import_periodic_auth_jwks" "thunder_import_periodic_auth_jwks" {
  auth_jwks     = "test"
  period        = 4699088
  use_mgmt_port = 1
  remote_file   = "test"
}