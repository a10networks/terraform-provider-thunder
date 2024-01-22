provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vpn_ipsec_group" "thunder_vpn_ipsec_group" {

  name = "test"
  ipsecgroup_cfg {
    ipsec    = "test"
    priority = 9
  }
  user_tag = "test"
}