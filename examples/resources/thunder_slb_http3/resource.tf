provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_http3" "thunder_slb_http3" {
  sampling_enable {
    counters1 = "all"
  }
}