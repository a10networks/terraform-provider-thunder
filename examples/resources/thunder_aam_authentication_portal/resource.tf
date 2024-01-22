provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_portal" "thunder_aam_authentication_portal" {
  name     = "default-portal"
  user_tag = "37"
}