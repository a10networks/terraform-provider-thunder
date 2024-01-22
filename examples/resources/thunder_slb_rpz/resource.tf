provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_rpz" "thunder_slb_rpz" {
  check = "37"
  sampling_enable {
    counters1 = "all"
  }
}