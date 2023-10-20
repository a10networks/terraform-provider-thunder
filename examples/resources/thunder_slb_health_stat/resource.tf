provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_health_stat" "test_thunder_slb_health_stat" {
  sampling_enable {
    counters1 = "all"
  }
}

