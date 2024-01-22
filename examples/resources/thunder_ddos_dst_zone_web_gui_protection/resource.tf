provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_web_gui_protection" "thunder_ddos_dst_zone_web_gui_protection" {
  zone_name = "testZone"
  ip_proto {
    proto_name_list {
      protocol = "icmp-v4"
      pbe      = "73"
      user_tag = "90"
    }
  }
  port {
    zone_service_list {
      port_num = 15010
      protocol = "dns-tcp"
      pbe      = "53"
    }
    zone_service_other_list {
      port_other = "other"
      protocol   = "tcp"
      pbe        = "50"
    }
  }
  port_range_list {
    port_range_start = 41654
    port_range_end   = 17241
    protocol         = "dns-tcp"
    pbe              = "26"
    user_tag         = "31"
  }

}