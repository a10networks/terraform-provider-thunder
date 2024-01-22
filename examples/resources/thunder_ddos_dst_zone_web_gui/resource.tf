provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_web_gui" "thunder_ddos_dst_zone_web_gui" {
  zone_name                = "testZone"
  activated_after_learning = 1
  create_time              = "9"
  learning {
    duration      = "6hour"
    starting_time = "8"
  }
  modify_time = "3"
  protection {
    port {
      zone_service_list {
        port_num = 38175
        protocol = "dns-tcp"
        pbe      = "116"
      }
      zone_service_other_list {
        port_other = "other"
        protocol   = "tcp"
        pbe        = "16"
      }
    }
    ip_proto {
      proto_name_list {
        protocol = "icmp-v4"
        pbe      = "57"
        user_tag = "6"
      }
    }
    port_range_list {
      port_range_start = 31609
      port_range_end   = 48578
      protocol         = "dns-tcp"
      pbe              = "117"
      user_tag         = "64"
    }
  }
  sensitivity = "3"
  status      = "newly"

}