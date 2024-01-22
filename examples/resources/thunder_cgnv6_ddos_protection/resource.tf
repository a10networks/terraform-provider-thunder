provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_ddos_protection" "thunder_cgnv6_ddos_protection" {
  disable_nat_ip_by_bgp {
  }
  enable_action = "local"
  ip_entries {
  }
  l4_entries {
  }
  logging_action = "enable"
  max_hw_entries = 262144
  sampling_enable {
    counters1 = "all"
  }
  syn_cookie {
    syn_cookie_enable       = 1
    syn_cookie_on_threshold = 17008
    syn_cookie_on_timeout   = 120
  }
  toggle = "enable"
  zone   = "8"
}