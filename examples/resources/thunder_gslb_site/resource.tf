provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_site" "thunder_gslb_site" {

  site_name = "3"
  auto_map  = 1
  disable   = 1
  weight    = 15
  multiple_geo_locations {
    geo_location = "68"
  }
  bw_cost    = 1
  limit      = 152721
  threshold  = 15
  controller = "125"
  user_tag   = "118"
  active_rdt {
    aging_time    = 123
    smooth_factor = 15
    range_factor  = 999
    limit         = 15272
    mask          = "/32"
    ipv6_mask     = 127
    ignore_count  = 11
    bind_geoloc   = 1
    overlap       = 1
  }
  ip_server_list {
    ip_server_name = "test-server1"
  }
  slb_dev_list {
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
        ipv4 = "10.10.10.11"
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
}