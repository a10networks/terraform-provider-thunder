provider "thunder" {
    address  = var.dut9049
    username = var.username
    password = var.password
}

resource "thunder_gslb_service_ip" "thunder_gslb_Service_ip_test" {
    node_name = "server_real1"
    ip_address = "2.2.2.2"
}
resource "thunder_gslb_site" "thunder_gslb_site_test" {
    site_name = "saleem"
    auto_map = 1
    weight = 100
    template = "snmp"
    active_rdt {
        aging_time = 20
        smooth_factor = 20
        range_factor = 30
        limit = 1638
        mask = "/24"
        ipv6_mask = "64"
        ignore_count = 10
        bind_geoloc = 1
        overlap = 1
    }
    ip_server_list {
        ip_server_name = "server_real"
    }
    slb_dev_list {
        device_name = "slb1"
        ip_address = "1.1.1.1"
        ipv6_address = "4001::1:10"
        admin_preference = 99
        proto_aging_time = 1000
        proto_compatible = 1
        msg_format_acos_2x = 1
        auto_detect = "ip"
        user_tag = "a10incedo"
        max_client = 1000
        vip_server{
            vip_server_name_list {
                vip_name = thunder_gslb_service_ip.thunder_gslb_Service_ip_test.node_name
            }
            vip_server_v4_list {
                 ipv4 = "1.1.1.1"
                 sampling_enable {
                counters1 = "all"
            }
            }
           vip_server_v6_list {
            ipv6 = "6001::1:20"
            sampling_enable {
                counters1 = "all"
            }
           } 
        }
    }
    multiple_geo_locations {
        geo_location = "North America,United States"
    }
}
