provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_fast_http_proxy" "test_thunder_slb_fast_http_proxy" {
  sampling_enable {
    counters1 = "all"
    counters2 = "req_sz_4k"
  }
}