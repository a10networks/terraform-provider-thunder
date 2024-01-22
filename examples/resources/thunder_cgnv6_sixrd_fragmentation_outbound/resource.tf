provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_sixrd_fragmentation_outbound" "thunder_cgnv6_sixrd_fragmentation_outbound" {

  action = "drop"
  df_set = "send-icmp"

}