provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_ip_nat_alg_pptp" "pptp1" {
  pptp = "enable"
  sampling_enable {
    counters1 = "all"
  }
}