provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_site_slb_dev" "thunder_gslb_site_slb_dev" {

  site_name          = "3"
  device_name        = "25"
  ip_address         = "10.10.10.10"
  admin_preference   = 100
  auto_detect        = "ip-and-port"
  auto_map           = 1
  max_client         = 32768
  proto_aging_time   = 32560
  proto_aging_fast   = 1
  gateway_ip_addr    = "10.10.10.10"
  proto_compatible   = 1
  msg_format_acos_2x = 1
  user_tag           = "109"
  vip_server {
    vip_server_v4_list {
      ipv4 = "10.1.20.11"
      sampling_enable {
        counters1 = "all"
      }
    }
    vip_server_v6_list {
      ipv6 = "cf58:9ec7:91f2:cd5c:219c:66eb:d05b:51ca"
      sampling_enable {
        counters1 = "all"
      }
    }
    vip_server_name_list {
      vip_name = "vs2"
      sampling_enable {
        counters1 = "all"
      }
    }
  }
}