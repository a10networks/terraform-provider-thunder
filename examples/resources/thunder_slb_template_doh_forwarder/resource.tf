provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_doh_forwarder" "thunder_slb_template_doh_forwarder" {

  name        = "test1"
  bypass_doh  = 1
  v4_internal = 0
  v4_l4_proto = "both"
  v4_port     = 53
  v6_internal = 0
  v6_l4_proto = "both"
  v6_port     = 53
}