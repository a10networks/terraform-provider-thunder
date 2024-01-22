provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_netflow_monitor" "thunder_netflow_monitor" {
  name = "a11"
  destination {
    ip_cfg {
      ip    = "10.0.2.7"
      port4 = 9996
    }
  }
  disable = 0
  disable_log_by_destination {
    icmp = 1
    ip_list {
      ipv4_addr = "10.10.10.10/24"
      tcp_list {
        tcp_port_start = 35324
        tcp_port_end   = 35324
      }
      udp_list {
        udp_port_start = 32422
        udp_port_end   = 32422
      }
      icmp     = 1
      others   = 1
      user_tag = "87"
    }
    ip6_list {
      ipv6_addr = "b184:baa0:dff5:52a2:78a2:7c9e:b1f0:4111/24"
      tcp_list {
        tcp_port_start = 31442
        tcp_port_end   = 31442
      }
      udp_list {
        udp_port_start = 23133
        udp_port_end   = 23133
      }
      icmp     = 1
      others   = 1
      user_tag = "19"
    }
    others = 0
    tcp_list {
      tcp_port_start = 43553
      tcp_port_end   = 43553
    }
    udp_list {
      udp_port_start = 32423
      udp_port_end   = 32423
    }
  }
  record {
    dslite            = 1
    nat44             = 1
    nat64             = 1
    netflow_v5        = 1
    netflow_v5_ext    = 1
    sesn_event_dslite = "both"
    sesn_event_fw4    = "both"
    sesn_event_fw6    = "both"
    sesn_event_nat44  = "both"
    sesn_event_nat64  = "both"
  }
  resend_template {
    records = 10011
    timeout = 18023
  }
  sample {
    ethernet_list {
      ifindex = 2
    }
  }
  sampling_enable {
    counters1 = "all"
  }
  source_address {
    ip = "10.0.2.5"
  }
  user_tag = "109"
}