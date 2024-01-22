provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_smpp" "thunder_slb_smpp" {
  sampling_enable {
    counters1 = "all"
  }
}