provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_generic_proxy" "generic_proxy" {
  sampling_enable {
    counters1 = "all"
  }
}
