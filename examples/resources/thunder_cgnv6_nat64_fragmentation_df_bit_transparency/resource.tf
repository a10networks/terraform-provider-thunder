provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat64_fragmentation_df_bit_transparency" "thunder_cgnv6_nat64_fragmentation_df_bit_transparency" {
  df_bit_value = "enable"
}