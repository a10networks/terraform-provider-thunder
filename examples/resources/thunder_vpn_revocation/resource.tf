provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vpn_revocation" "thunder_vpn_revocation" {
  name     = "test"
  user_tag = "test_vpn_revocation"
}