provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_global" "thunder_cgnv6_lsn_global" {
  attempt_port_preservation = "disable"
  enhanced_user_tracking    = 0
  hairpinning               = "filter-none"
  half_close_timeout        = 1470
  icmp {
    send_on_port_unavailable    = "disable"
    send_on_user_quota_exceeded = "admin-filtered"
  }
  inbound_refresh           = "enable"
  inbound_refresh_full_cone = "enable"
  ip_selection              = "random"
  port_batching {
    size                   = "1"
    tcp_time_wait_interval = 2
  }
  sampling_enable {
    counters1 = "all"
    counters2 = "user_quota_not_found"
    counters3 = "specific_port_allocated_from_quota_partition"
    counters4 = "acl_http_domain_node_exceeded"
  }
  strictly_sticky_nat = 0
  syn_timeout         = 4
}