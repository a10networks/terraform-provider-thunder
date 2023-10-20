provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_l4" "test_thunder_slb_l4" {
  sampling_enable {
    counters1 = "all"
  }
}

