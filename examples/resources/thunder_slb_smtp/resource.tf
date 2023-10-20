provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_smtp" "test_thunder_slb_smtp" {
  sampling_enable {
    counters1 = "all"
  }
}

