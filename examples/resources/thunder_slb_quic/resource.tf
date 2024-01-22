provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_quic" "thunder_slb_quic" {
  sampling_enable {
    counters1 = "all"
    counters2 = "Err_frame_dec1"
  }
}