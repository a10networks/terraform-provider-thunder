provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_ftp_ctl" "test_thunder_slb_ftp_ctl" {
  sampling_enable {
    counters1 = "all"
  }
}