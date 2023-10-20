provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_health_gateway" "health_gateway" {
  sampling_enable {
    counters1 = "all"
  }
}
