provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lw_4o6_global" "thunder_cgnv6_lw_4o6_global" {

  hairpinning = "filter-all"
  no_forward_match {
    send_icmpv6 = 1
  }
  no_reverse_match {
    send_icmp = 1
  }
  sampling_enable {
    counters1 = "all"
  }
}