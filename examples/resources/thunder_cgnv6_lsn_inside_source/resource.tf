provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_inside_source" "thunder_cgnv6_lsn_inside_source" {

  class_list = "31"

}