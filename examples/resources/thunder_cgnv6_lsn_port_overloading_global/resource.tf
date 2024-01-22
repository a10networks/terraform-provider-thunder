provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_port_overloading_global" "thunder_cgnv6_lsn_port_overloading_global" {

  allow_different_user = 1
  unique               = "destination-address-and-port"

}