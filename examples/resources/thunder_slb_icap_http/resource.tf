provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_icap_http" "thunder_slb_icap_http" {
  sampling_enable {
    counters1 = "all"
  }
}