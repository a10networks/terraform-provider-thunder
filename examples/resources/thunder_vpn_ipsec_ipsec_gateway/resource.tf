provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vpn_ipsec_ipsec_gateway" "thunder_vpn_ipsec_ipsec_gateway" {
  ike_gateway = "test"
  name        = "test"

}