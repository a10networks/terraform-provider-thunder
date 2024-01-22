provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_delete_auth_jwks" "thunder_delete_auth_jwks" {
  jwk_name = "test"
}