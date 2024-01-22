provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat64_fragmentation_inbound" "thunder_cgnv6_nat64_fragmentation_inbound" {

  count1      = 1
  df_set      = "drop"
  frag_action = "ipv6"

}