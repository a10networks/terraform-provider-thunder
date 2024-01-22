provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lw_4o6_fragmentation_outbound" "thunder_cgnv6_lw_4o6_fragmentation_outbound" {
  df_set      = "drop"
  frag_action = "ipv4"
}