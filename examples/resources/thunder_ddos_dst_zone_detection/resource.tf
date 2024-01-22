provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_detection" "thunder_ddos_dst_zone_detection" {
  zone_name = "testZone"
  notification {
    configuration = "configuration"
  }
  outbound_detection {
    configuration = "configuration"
    toggle        = "disable"
    indicator_list {
      type          = "pkt-rate"
      threshold_num = 2096039599
      user_tag      = "17"
    }
  }
  packet_anomaly_detection {
    configuration = "configuration"
    toggle        = "enable"
    indicator_list {
      type          = "port-zero-pkt-rate"
      threshold_num = 100
      user_tag      = "111"
    }
  }
  service_discovery {
    configuration      = "configuration"
    toggle             = "disable"
    pkt_rate_threshold = 10
  }
  settings = "settings"
  toggle   = "enable"
  victim_ip_detection {
    configuration    = "configuration"
    toggle           = "disable"
    histogram_toggle = "histogram-disable"
    indicator_list {
      type             = "pkt-rate"
      ip_threshold_num = 2036967623
      user_tag         = "99"
    }
  }

}