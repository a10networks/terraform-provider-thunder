provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_common_dns_response_rate_limiting" "thunder_slb_common_dns_response_rate_limiting" {
  max_table_entries = 2088558
}