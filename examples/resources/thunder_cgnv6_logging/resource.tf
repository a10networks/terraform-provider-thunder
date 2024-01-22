provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_logging" "thunder_cgnv6_logging" {
  nat_quota_exceeded {
    level = "warning"
  }
  nat_resource_exhausted {
    level = "critical"
  }
  sampling_enable {
    counters1 = "all"
  }
  source_address {
  }
  tcp_svr_status {
  }
}