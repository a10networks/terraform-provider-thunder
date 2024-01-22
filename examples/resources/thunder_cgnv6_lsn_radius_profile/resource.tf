provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_radius_profile" "thunder_cgnv6_lsn_radius_profile" {
  lid_profile_index = 9
  radius {
    attribute           = "custom1"
    starts_with         = "49"
    starts_with_lsn_lid = 619
  }
  user_tag = "88"
}