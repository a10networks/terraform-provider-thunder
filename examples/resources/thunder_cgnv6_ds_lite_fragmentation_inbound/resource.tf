provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_ds_lite_fragmentation_inbound" "thunder_cgnv6_ds_lite_fragmentation_inbound" {
  count1      = 1
  df_set      = "ipv6"
  frag_action = "ipv4"
}