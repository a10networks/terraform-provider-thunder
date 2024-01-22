provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_sip" "thunder_slb_sip" {
  sampling_enable {
    counters1 = "all"
  }
}