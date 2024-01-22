provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_flowspec_fragmentation_option" "thunder_flowspec_fragmentation_option" {

  name           = "test"
  frag_attribute = "is-fragment"
}