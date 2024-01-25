provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_rule_set_rule_stats" "thunder_rule_set_rule_stats" {
  stats {
    hit_count            = hit_count
    permit_bytes         = permit_bytes
    deny_bytes           = deny_bytes
    reset_bytes          = reset_bytes
    permit_packets       = permit_packets
    deny_packets         = deny_packets
    reset_packets        = reset_packets
    active_session_tcp   = active_session_tcp
    active_session_udp   = active_session_udp
    active_session_icmp  = active_session_icmp
    active_session_other = active_session_other
    session_tcp          = session_tcp
    session_udp          = session_udp
    session_icmp         = session_icmp
    session_other        = session_other
    active_session_sctp  = active_session_sctp
    session_sctp         = session_sctp
    hitcount_timestamp   = hitcount_timestamp
    rate_limit_drops     = rate_limit_drops
  }

}
output "get_rule_set_rule_stats" {
  value = ["${data.thunder_rule_set_rule_stats.thunder_rule_set_rule_stats}"]
}