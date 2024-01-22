provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_lid" "thunder_cgnv6_lsn_lid" {
  lid_number          = 635
  name                = "test"
  override            = "none"
  respond_to_user_mac = 0
}