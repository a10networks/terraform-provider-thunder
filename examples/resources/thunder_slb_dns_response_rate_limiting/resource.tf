provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_dns_response_rate_limiting" "test_thunder_slb_dns_response_rate_limiting" {
  sampling_enable {
    counters1 = "all"
  }
}