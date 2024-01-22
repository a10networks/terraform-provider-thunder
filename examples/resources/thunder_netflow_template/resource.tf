provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_netflow_template" "thunder_netflow_template" {
  name = 2781
  information_element_blk {
    information_element = "fwd-tuple-vnp-id"
  }
  user_tag = "54"
}