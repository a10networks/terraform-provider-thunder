provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_persist" "thunder_slb_persist" {
  sampling_enable {
    counters1 = "all"
  }
}