provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_ipv6_nat_pool" "A" {
  pool_name     = "K"
  start_address = "a::1"
  end_address   = "a::1"
  netmask       = 98
  sampling_enable {
    counters1 = "all"
  }
}