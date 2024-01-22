provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_ospf" "thunder_router_ospf" {
  process_id                    = 400
  auto_cost_reference_bandwidth = 30000
  bfd_all_interfaces            = 1
  rfc1583_compatible            = 1
  default_metric                = 30

  distance {
    distance_value = 10
    distance_ospf {
      distance_external   = 20
      distance_inter_area = 30
      distance_intra_area = 40
    }
  }
  distribute_internal_list {
    di_area_ipv4 = "1.1.1.1"
    di_area_num  = 200
    di_cost      = 3000
    di_type      = "lw4o6"
  }
  ha_standby_extra_cost {
    extra_cost = 500
    group      = 2
  }
  router_id {
    value = "5.5.5.5"
  }
  ospf_1 {
    abr_type {
      option = "standard"
    }
  }
  log_adjacency_changes_cfg {
    state = "detail"
  }
  max_concurrent_dd = 10
  maximum_area      = 2000
  user_tag          = "ospf"
  neighbor_list {
    address = "1.1.1.1"
    cost    = 400
  }
  distribute_lists {
    direction = "in"
    value     = "ACL1"
  }
  host_list {
    area_cfg {
      area_ipv4 = "0.0.0.0"
      area_num  = 0
      cost      = 65535
    }
    host_address = "2.2.2.2"
  }
  overflow {
    database {
      limit = "hard"
    }
  }
  passive_interface {
    eth_cfg {
      ethernet    = 2
      eth_address = "2.2.2.2"
    }
    lif_cfg {
      lif         = 3
      lif_address = "4.4.4.4"
    }
    ve_cfg {
      ve         = 300
      ve_address = "5.5.5.5"
    }
    trunk_cfg {
      trunk         = 5
      trunk_address = "6.6.6.6"
    }
    tunnel_cfg {
      tunnel         = 20
      tunnel_address = "12.12.12.13"
    }
    loopback_cfg {
      loopback         = 1
      loopback_address = "11.11.11.11"
    }

  }
  summary_address_list {
    summary_address = "20.20.20.0/24"
    tag             = 3000
    not_advertise   = 0
  }
  timers {
    spf {
      exp {
        min_delay = 3000
        max_delay = 40000
      }
    }
  }
  network_list {
    network_ipv4      = "1.1.1.1"
    network_ipv4_mask = "255.255.255.0"
    network_area {
      network_area_ipv4 = "0.0.0.0"
      network_area_num  = 0
      instance_value    = 0
    }
  }
  default_information {
    originate   = 1
    always      = 1
    metric      = 40
    metric_type = 1
    route_map   = "MAP"
  }
  redistribute {
    ip_nat             = 1
    metric_ip_nat      = 40
    metric_type_ip_nat = 1
    route_map_ip_nat   = "NAT"
    tag_ip_nat         = 800
    ip_nat_floating_list {
      ip_nat_prefix              = "9.9.9.9/24"
      ip_nat_floating_ip_forward = "2.2.2.2"
    }
    vip_list {
      type_vip        = "only-flagged"
      metric_type_vip = 1
      metric_vip      = 800
      route_map_vip   = "VIP"
      tag_vip         = 500
    }
    vip_floating_list {
      vip_address             = "5.5.5.5"
      vip_floating_ip_forward = "6.6.6.6"
    }
    ospf_list {
      process_id       = 200
      ospf             = 1
      metric_ospf      = 50
      metric_type_ospf = 1
      route_map_ospf   = "OSPF"
      tag_ospf         = 500
    }
  }
  area_list {
    area_ipv4 = "0.0.0.2"
    area_num  = 0
    auth_cfg {
      authentication = 1
      message_digest = 1
    }
    filter_lists {
      filter_list     = 1
      acl_name        = "ACL1"
      acl_direction   = "in"
      plist_name      = "PREFIX1"
      plist_direction = "in"
    }
    default_cost = 1

    range_list {
      area_range_prefix = "1.1.1.1/24"
      option            = "advertise"
    }
    virtual_link_list {
      virtual_link_ip_addr        = "3.3.3.3"
      bfd                         = 1
      hello_interval              = 10
      dead_interval               = 40
      retransmit_interval         = 300
      transmit_delay              = 50
      virtual_link_authentication = 1
      virtual_link_auth_type      = "message-digest"
      authentication_key          = "string"
      message_digest_key          = 20
      md5                         = "MD5"
    }
  }

}