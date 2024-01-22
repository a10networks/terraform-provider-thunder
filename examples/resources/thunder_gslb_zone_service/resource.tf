provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_zone_service" "thunder_gslb_zone_service" {

  name         = "a11"
  service_name = "s25"
  service_port = 33159
  action       = "forward"
  forward_type = "both"
  disable      = 0
  dns_a_record {
    dns_a_record_srv_list {
      svrname    = "test-server1"
      no_resp    = 0
      as_backup  = 0
      weight     = 96
      ttl        = 0
      as_replace = 0
      disable    = 1
      static     = 1
      admin_ip   = 198
    }
    dns_a_record_ipv4_list {
      dns_a_record_ip = "10.10.10.10"
      no_resp         = 0
      as_backup       = 0
      weight          = 31
      ttl             = 0
      as_replace      = 0
      disable         = 0
      static          = 1
      admin_ip        = 232
    }
    dns_a_record_ipv6_list {
      dns_a_record_ipv6 = "2001:db8:3333:4444:5555:6666:7777:8888"
      no_resp           = 0
      as_backup         = 0
      weight            = 97
      ttl               = 0
      as_replace        = 0
      disable           = 0
      static            = 1
      admin_ip          = 202
    }
  }
  dns_caa_record_list {
    critical_flag = 106
    property_tag  = "32"
    rdata         = "798"
    ttl           = 0
    sampling_enable {
      counters1 = "all"
    }
  }
  dns_cname_record_list {
    alias_name       = "89"
    admin_preference = 100
    weight           = 1
    as_backup        = 0
    sampling_enable {
      counters1 = "all"
    }
  }
  dns_mx_record_list {
    mx_name  = "112"
    priority = 35640
    ttl      = 0
    sampling_enable {
      counters1 = "all"
    }
  }
  dns_naptr_record_list {
    naptr_target  = "35"
    service_proto = "68"
    flag          = "1"
    order         = 8489
    preference    = 16018
    regexp        = 0
    ttl           = 0
    sampling_enable {
      counters1 = "all"
    }
  }
  dns_ns_record_list {
    ns_name = "12"
    ttl     = 0
    sampling_enable {
      counters1 = "all"
    }
  }
  dns_ptr_record_list {
    ptr_name = "125"
    ttl      = 0
    sampling_enable {
      counters1 = "all"
    }
  }
  dns_record_list {
    type = 57856
    data = "438"
  }
  dns_srv_record_list {
    srv_name = "70"
    port     = 27600
    priority = 58080
    weight   = 10
    ttl      = 0
    sampling_enable {
      counters1 = "all"
    }
  }
  dns_txt_record_list {
    record_name = "11"
    txt_data    = "126"
    ttl         = 0
    sampling_enable {
      counters1 = "all"
    }
  }
  geo_location_list {
    geo_name = "75"
    user_tag = "70"
  }
  health_check_gateway = "enable"
  health_check_port {
    health_check_port = 37739
  }
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "56"
}