provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_map_encapsulation_fragmentation_inbound" "thunder_cgnv6_map_encapsulation_fragmentation_inbound" {

  df_set      = "drop"
  frag_action = "ipv6"

}