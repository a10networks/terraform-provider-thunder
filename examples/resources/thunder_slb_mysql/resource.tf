provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_mysql" "test_thunder_slb_mysql" {
  sampling_enable {
    counters1 = "all"
  }
}

