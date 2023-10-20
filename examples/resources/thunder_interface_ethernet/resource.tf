provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_interface_ethernet" "IntEthernet" {
  ifnum                = 3
  action               = "enable"
  name                 = "int"
  mtu                  = 1500
  l3_vlan_fwd_disable  = 1
  load_interval        = 10
  ping_sweep_detection = "enable"
  port_scan_detection  = "enable"

  icmp_rate_limit {
    lockup        = 800
    lockup_period = 1000
    normal        = 600
  }

  icmpv6_rate_limit {
    lockup_v6        = 600
    lockup_period_v6 = 1000
    normal_v6        = 200
  }
  access_list {
    acl_id = thunder_access_list_standard.resourceAccessListStandardTest.std
  }
  sampling_enable {
    counters1 = "all"
  }
  map {
    inside        = 1
    outside       = 0
    map_t_inside  = 0
    map_t_outside = 0
  }
  lw_4o6 {
    inside  = 1
    outside = 0
  }

  bfd {
    authentication {
      key_id   = 5
      method   = "md5"
      password = "bfd"
    }
    echo   = 0
    demand = 0
    interval_cfg {
      interval   = 200
      min_rx     = 200
      multiplier = 3
    }
  }

  isis {
    bfd_cfg {
      bfd     = 1
      disable = 0
    }
    circuit_type        = "level-1"
    padding             = 1
    lsp_interval        = 43
    network             = "broadcast"
    retransmit_interval = 30
    csnp_interval_list {
      csnp_interval = 20
      level         = "level-1"
    }
    priority_list {
      priority = 100
    }
    hello_interval_minimal_list {
      hello_interval_minimal = 1
      level                  = "level-1"
    }

    mesh_group {
      blocked = 0
    }

    wide_metric_list {
      wide_metric = 64
      level       = "level-1"
    }
    metric_list {
      metric = 30
      level  = "level-1"
    }
    hello_multiplier_list {
      hello_multiplier = 2
      level            = "level-1"
    }
    authentication {
      send_only_list {
        level     = "level-1"
        send_only = 1
      }
      key_chain_list {
        key_chain = "md5"
        level     = "level-1"
      }
      mode_list {
        level = "level-1"
        mode  = "md5"
      }

    }

  }
  ipv6 {
    address_list {
      ipv6_addr    = "3000::1/64"
      address_type = "anycast"
    }
    router_adver {
      action                = "enable"
      hop_limit             = 230
      max_interval          = 500
      min_interval          = 100
      default_lifetime      = 1700
      rate_limit            = 100000
      reachable_time        = 1
      retransmit_timer      = 1
      adver_mtu_disable     = 1
      managed_config_action = "enable"
      other_config_action   = "disable"
      adver_vrid_default    = 0

      prefix_list {
        prefix             = "1:1:2:2:3::1/64"
        not_on_link        = 0
        not_autonomous     = 0
        preferred_lifetime = 0
        valid_lifetime     = 0
      }

    }

    router {

      ripng {
        rip = 1
      }
      ospf {
        area_list {
          area_id_addr = "10.10.10.0"
          area_id_num  = 3
          tag          = "V3"
          instance_id  = 11
        }
      }

    }
    rip {
      split_horizon_cfg {
        state = "enable"
      }
    }

    ospf {
      bfd     = 0
      disable = 0
      cost_cfg {
        cost        = 120
        instance_id = 50
      }
      hello_interval_cfg {
        hello_interval = 15
        instance_id    = 50
      }
      priority_cfg {
        priority = 150
      }
      mtu_ignore_cfg {
        instance_id = 50
        mtu_ignore  = 1
      }
      retransmit_interval_cfg {
        retransmit_interval = 100
        instance_id         = 50
      }
      neighbor_cfg {
        neig_inst              = 20
        neighbor               = "6565::1"
        neighbor_cost          = 200
        neighbor_priority      = 150
        neighbor_poll_interval = 4567
      }
      network_list {
        broadcast_type      = "broadcast"
        network_instance_id = 50
        p2mp_nbma           = 0
      }
      transmit_delay_cfg {
        instance_id    = 50
        transmit_delay = 40
      }
      dead_interval_cfg {
        dead_interval = 60
        instance_id   = 50
      }
    }
  }
  lldp {
    notification_cfg {
      notif_enable = 1
      notification = 1
    }
    enable_cfg {
      tx        = 1
      rx        = 1
      rt_enable = 1
    }
    tx_dot1_cfg {
      link_aggregation = 1
      vlan             = 1
      tx_dot1_tlvs     = 1
    }
    tx_tlvs_cfg {
      tx_tlvs             = 1
      port_description    = 1
      exclude             = 1
      management_address  = 1
      system_name         = 1
      system_capabilities = 1
      system_description  = 1
    }
  }
  ip {
    allow_promiscuous_vip     = 1
    syn_cookie                = 1
    cache_spoofing_port       = 1
    ttl_ignore                = 1
    server                    = 1
    inside                    = 1
    outside                   = 0
    client                    = 0
    generate_membership_query = 1
    slb_partition_redirect    = 1
    helper_address_list {
      helper_address = "6.6.6.6"
    }
    address_list {
      ipv4_address = "2.2.2.2"
      ipv4_netmask = "255.255.0.0"
    }

    router {
      isis {
        tag = "isis"
      }
    }
    rip {
      authentication {

        mode {
          mode = "md5"
        }
        key_chain {
          key_chain = "ripat"
        }
      }
      send_packet    = 1
      receive_packet = 1
      receive_cfg {
        receive = 1
        version = 2
      }
      send_cfg {
        send    = 1
        version = 2
      }
      split_horizon_cfg {
        state = "enable"
      }
    }

    ospf {
      ospf_ip_list {
        ip_addr            = "33.5.4.4"
        authentication     = 1
        value              = "null"
        authentication_key = "oppsf"
        cost               = 500
        database_filter    = "all"
        out                = 1
        dead_interval      = 40
        hello_interval     = 10
        message_digest_cfg {
          message_digest_key = 4
          md5_value          = "md5"

        }
        mtu_ignore          = 0
        priority            = 120
        retransmit_interval = 50
        transmit_delay      = 100
      }
      ospf_global {
        cost                = 100
        mtu_ignore          = 1
        retransmit_interval = 50
        dead_interval       = 80
        hello_interval      = 10
        priority            = 140
        network {
          broadcast           = 1
          non_broadcast       = 0
          point_to_point      = 0
          point_to_multipoint = 0
          p2mp_nbma           = 0

        }
        mtu            = 1400
        transmit_delay = 70
        disable        = "all"
        message_digest_cfg {
          message_digest_key = 5
          md5_value          = "md5"
        }
        authentication_cfg {
          value          = "null"
          authentication = 1
        }
        database_filter_cfg {
          database_filter = "all"
          out             = 1
        }
        bfd_cfg {
          bfd     = 1
          disable = 0
        }
      }
    }
  }
}