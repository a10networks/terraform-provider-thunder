provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_vrid_lead" "thunder_vrrp_a_vrid_lead" {
  user_tag      = "test"
  vrid_lead_str = "test"
}