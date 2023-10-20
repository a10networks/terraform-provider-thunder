provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_mlb" "test_thunder_slb_mlb" {
  sampling_enable {
    counters1 = "all"
  }
}

