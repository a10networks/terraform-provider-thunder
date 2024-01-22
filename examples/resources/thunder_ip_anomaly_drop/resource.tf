provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_anomaly_drop" "thunder_ip_anomaly_drop" {
  bad_content = 27
  drop_all    = 1
  frag        = 1
  ip_option   = 1
  ipv6_ext_header {
    ipv6_eh_frag      = 1
    ipv6_eh_auth      = 1
    ipv6_eh_esp       = 1
    ipv6_eh_mobility  = 1
    ipv6_eh_nonext    = 1
    ipv6_eh_malformed = 1
    ipv6_eh_hbh       = 1
    ipv6_eh_dest      = 1
    ipv6_eh_routing   = 1
    hbh_option_list {
      hbh_otype_from = 233
      hbh_otype_to   = 233
    }
    dst_option_list {
      dst_otype_from = 211
      dst_otype_to   = 211
    }
    routing_option_list {
      routing_otype_from = 121
      routing_otype_to   = 121
    }
    unknown_ext_header_list {
      eh_type_from = 251
      eh_type_to   = 251
    }
  }
  land_attack     = 1
  out_of_sequence = 126
  packet_deformity {
    packet_deformity_layer_3 = 1
    packet_deformity_layer_4 = 1
  }
  ping_of_death = 1
  sampling_enable {
    counters1 = "all"
  }
  security_attack {
    security_attack_layer_3 = 1
    security_attack_layer_4 = 0
  }
  tcp_no_flag  = 1
  tcp_syn_fin  = 1
  tcp_syn_frag = 1
  zero_window  = 124
}