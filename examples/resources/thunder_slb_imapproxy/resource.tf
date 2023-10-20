provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_imapproxy" "test_thunder_slb_imapproxy" {
  sampling_enable {
    counters1 = "all"
  }
}

