provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat46_stateless_fragmentation_outbound" "thunder_cgnv6_nat46_stateless_fragmentation_outbound" {

  action = "ipv6"
  count1 = 1
  df_set = "drop"

}