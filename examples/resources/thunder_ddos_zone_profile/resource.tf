provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_profile" "thunder_ddos_zone_profile" {
  profile_name = "test"
  user_tag     = "test"
  port_list {
    port_num      = 22
    port_protocol = "dns-tcp"
    user_tag      = "test"
    indicator_list {
      indicator_name = "pkt-drop-rate"
      src_threshold_cfg {
        src_threshold_num = 3
      }
      zone_threshold_cfg {
        zone_threshold_num = 4
      }

      user_tag = "test"
    }
  }
  ip_proto {
    proto_name_list {
      protocol = "icmp-v4"
      indicator_list {
        indicator_name = "pkt-rate"
        src_threshold_cfg {
          src_threshold_num = 22
        }
        zone_threshold_cfg {
          zone_threshold_num = 3
        }
        user_tag = "test"
      }
    }
  }
  port_range_list {
    port_range_start = 3
    port_range_end   = 4
    protocol         = "udp"
    user_tag         = "test"
    indicator_list {
      indicator_name = "pkt-rate"
      src_threshold_cfg {
        src_threshold_num = 6
      }
      zone_threshold_cfg {
        zone_threshold_num = 3
      }
      user_tag = "test"
    }
  }
}