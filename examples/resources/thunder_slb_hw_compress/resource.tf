provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_hw_compress" "test_thunder_slb_hw_compress" {
  sampling_enable {
    counters1 = "all"
  }
}

