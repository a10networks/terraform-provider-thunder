provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_gslb_zone" "thunder_gslb_zone_test" {
  name = "z1"
  service_list {
    service_port = 80
    service_name = "http"
    action       = "drop"
    dns_mx_record_list {
      mx_name  = "saleem"
      priority = 1000
      ttl      = 1000
      sampling_enable {
        counters1 = "all"
      }
    }
    dns_ns_record_list {
      ns_name = "ns_name"
      ttl     = 100
      sampling_enable {
        counters1 = "all"
      }
    }
    dns_ptr_record_list {
      ptr_name = "ptr_name"
      ttl      = 100
      sampling_enable {
        counters1 = "all"
      }
    }
    dns_a_record {
      dns_a_record_ipv4_list {
        dns_a_record_ip = "10.10.2.1"
        as_backup       = 1
        no_resp         = 1
        weight          = 100
        ttl             = 10
        as_replace      = 1
      }
      dns_a_record_srv_list {
        svrname    = "svrname"
        no_resp    = 1
        as_backup  = 1
        weight     = 100
        ttl        = 10
        as_replace = 1
      }
      dns_a_record_ipv6_list {
        dns_a_record_ipv6 = "4001::1:10"
        as_backup         = 1
        no_resp           = 1
        weight            = 100
        ttl               = 10
        as_replace        = 1
      }
    }
  }
}