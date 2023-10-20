provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_ftp_data" "test_thunder_slb_ftp_data" {
  sampling_enable {
    counters1 = "all"
  }
}

