provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_zone" "thunder_gslb_zone" {

  name    = "a11"
  disable = 1
  dns_caa_record_list {
    critical_flag = 181
    property_tag  = "critical"
    rdata         = "103"
    sampling_enable {
      counters1 = "all"
    }
  }
  dns_mx_record_list {
    mx_name  = "mx79"
    priority = 6888
    sampling_enable {
      counters1 = "all"
    }
  }
  dns_ns_record_list {
    ns_name = "89"
    ttl     = 662713792
    sampling_enable {
      counters1 = "all"
    }
  }
  dns_soa_record {
    soa_name   = "61"
    mail       = "50"
    expire     = 1209600
    refresh    = 3600
    retry      = 900
    serial     = 18717180
    soa_ttl    = 21052207
    external   = "s1"
    ex_mail    = "58"
    ex_expire  = 1209600
    ex_refresh = 3600
    ex_retry   = 900
    ex_serial  = 85660141
    ex_soa_ttl = 13197458
  }
  sampling_enable {
    counters1 = "all"
  }
  service_list {
    service_port         = 80
    service_name         = "s38"
    action               = "forward"
    forward_type         = "both"
    disable              = 0
    health_check_gateway = "enable"
    health_check_port {
      health_check_port = 12635
    }
    user_tag = "25"
    sampling_enable {
      counters1 = "all"
    }
    dns_a_record {
      dns_a_record_srv_list {
        svrname    = "test-server1"
        no_resp    = 1
        as_backup  = 1
        weight     = 100
        ttl        = 352732
        as_replace = 1
        disable    = 1
        static     = 1
        admin_ip   = 56
      }
      dns_a_record_ipv4_list {
        dns_a_record_ip = "10.10.10.10"
        no_resp         = 0
        as_backup       = 1
        weight          = 99
        ttl             = 0
        as_replace      = 0
        disable         = 0
        static          = 1
        admin_ip        = 1
      }
      dns_a_record_ipv6_list {
        dns_a_record_ipv6 = "2001:db8:3333:4444:5555:6666:7777:8888"
        no_resp           = 0
        as_backup         = 1
        weight            = 65
        ttl               = 0
        as_replace        = 0
        disable           = 0
        static            = 1
        admin_ip          = 181
      }
    }
    dns_cname_record_list {
      alias_name       = "89"
      admin_preference = 101
      weight           = 99
      as_backup        = 0
      sampling_enable {
        counters1 = "all"
      }
    }
    dns_mx_record_list {
      mx_name  = "26"
      priority = 57954
      sampling_enable {
        counters1 = "all"
      }
    }
    dns_ns_record_list {
      ns_name = "54"
      ttl     = 627424
      sampling_enable {
        counters1 = "all"
      }
    }
    dns_ptr_record_list {
      ptr_name = "67"
      ttl      = 573187
      sampling_enable {
        counters1 = "all"
      }
    }
    dns_srv_record_list {
      srv_name = "90"
      port     = 64338
      priority = 6644
      weight   = 10
      sampling_enable {
        counters1 = "all"
      }
    }
    dns_naptr_record_list {
      naptr_target  = "30"
      service_proto = "34"
      flag          = "1"
      order         = 5352
      preference    = 35430
      sampling_enable {
        counters1 = "all"
      }
    }
    dns_txt_record_list {
      record_name = "21"
      txt_data    = "204"
      ttl         = 45744
      sampling_enable {
        counters1 = "all"
      }
    }
    dns_caa_record_list {
      critical_flag = 81
      property_tag  = "critical"
      rdata         = "3296"
      ttl           = 53733
      sampling_enable {
        counters1 = "all"
      }
    }
    dns_record_list {
      type = 13783
      data = "316"
    }
    geo_location_list {
      geo_name     = "20"
      action       = 1
      action_type  = "forward"
      forward_type = "both"
      user_tag     = "39"
    }
  }
  ttl      = 3839239
  user_tag = "98"
}