provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_icap" "thunder_slb_icap" {
  sampling_enable {
    counters1 = "all"
  }
}