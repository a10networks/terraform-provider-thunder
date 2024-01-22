provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_secure_gaming" "thunder_slb_secure_gaming" {
  sampling_enable {
    counters1 = "all"
  }
}