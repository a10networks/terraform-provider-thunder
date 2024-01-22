provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_http_alg" "thunder_cgnv6_http_alg" {
  sampling_enable {
    counters1 = "all"
  }
}