provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_topn_gtp_network_element_topn_node" "thunder_visibility_topn_gtp_network_element_topn_node" {

  activate = "test"
}