provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_ethernet" "thunder_interface_ethernet" {
  access_list {
    acl_id = 155
  }
  action          = "disable"
  auto_neg_enable = 0
  bfd {
    authentication {
      key_id   = 114
      method   = "md5"
      password = "6"
    }
    echo   = 0
    demand = 0
    interval_cfg {
      interval   = 494
      min_rx     = 363
      multiplier = 32
    }
  }
  cpu_process     = 0
  cpu_process_dir = "primary"
  ddos {
    outside = 0
    inside  = 0
  }
  duplexity                  = "auto"
  fec_forced_off             = 0
  flow_control               = 0
  gaming_protocol_compliance = 0
  icmp_rate_limit {
    normal        = 17778
    lockup        = 60272
    lockup_period = 1279
  }
  icmpv6_rate_limit {
    normal_v6        = 763
    lockup_v6        = 27440
    lockup_period_v6 = 15322
  }
  ifnum = ifnum
  ip {
    dhcp = 0
    address_list {
    }
    allow_promiscuous_vip = 0
    cache_spoofing_port   = 0
    helper_address_list {
      helper_address = "10.10.10.10"
    }
    inside                    = 0
    outside                   = 0
    ttl_ignore                = 0
    syn_cookie                = 0
    slb_partition_redirect    = 0
    generate_membership_query = 0
    query_interval            = 125
    max_resp_time             = 100
    client                    = 0
    unnumbered                = 0
    stateful_firewall {
      inside      = 0
      class_list  = "63"
      outside     = 0
      access_list = 0
      acl_id      = 106
    }
    router {
      isis {
        tag = "79"
      }
    }
    rip {
      authentication {
        str {
          string = "8"
        }
        mode {
          mode = "text"
        }
        key_chain {
        }
      }
      send_packet    = 1
      receive_packet = 1
      send_cfg {
        send    = 0
        version = "1"
      }
      receive_cfg {
        receive = 0
        version = "1"
      }
      split_horizon_cfg {
        state = "poisoned"
      }
    }
    ospf {
      ospf_global {
        authentication_cfg {
          authentication = 0
        }
        authentication_key = "2"
        bfd_cfg {
          bfd     = 0
          disable = 0
        }
        cost = 42828
        database_filter_cfg {
          database_filter = "all"
          out             = 0
        }
        dead_interval  = 40
        disable        = "all"
        hello_interval = 10
        message_digest_cfg {
          message_digest_key = 158
          md5 {
            md5_value = "13"
          }
        }
        mtu        = 48727
        mtu_ignore = 0
        network {
          broadcast = 0
          p2mp_nbma = 0
        }
        priority            = 1
        retransmit_interval = 5
        transmit_delay      = 1
      }
      ospf_ip_list {
        ip_addr            = "10.10.10.10"
        authentication     = 0
        authentication_key = "8"
        cost               = 60186
        database_filter    = "all"
        out                = 0
        dead_interval      = 40
        hello_interval     = 10
        message_digest_cfg {
          message_digest_key = 210
          md5_value          = "15"
        }
        mtu_ignore          = 0
        priority            = 1
        retransmit_interval = 5
        transmit_delay      = 1
      }
    }
  }
  ipg_bit_time = 96
  ipv6 {
    address_list {
      address_type = "anycast"
    }
    inside      = 0
    outside     = 0
    ipv6_enable = 0
    ttl_ignore  = 0
    access_list_cfg {
      v6_acl_name = "2"
      inbound     = 0
    }
    router_adver {
      action            = "disable"
      hop_limit         = 255
      max_interval      = 600
      min_interval      = 200
      default_lifetime  = 1800
      rate_limit        = 100000
      reachable_time    = 0
      retransmit_timer  = 0
      adver_mtu_disable = 1
      prefix_list {
        prefix             = "2001:db8:3333:4444:5555:6666:7777:8888"
        not_autonomous     = 0
        not_on_link        = 0
        preferred_lifetime = 604800
        valid_lifetime     = 2592000
      }
      managed_config_action        = "disable"
      other_config_action          = "disable"
      adver_vrid                   = 16
      use_floating_ip              = 0
      floating_ip                  = "2001:db8:3333:4444:5555:6666:7777:8888"
      use_floating_ip_default_vrid = 0
      floating_ip_default_vrid     = "2001:db8:3333:4444:5555:6666:7777:8888"
    }
    stateful_firewall {
      inside      = 0
      class_list  = "53"
      outside     = 0
      access_list = 0
    }
    router {
      ripng {
        rip = 0
      }
      ospf {
        area_list {
          area_id_num  = 3393471670
          area_id_addr = "10.10.10.10"
          tag          = "38"
          instance_id  = 0
        }
      }
      isis {
        tag = "123"
      }
    }
    rip {
      split_horizon_cfg {
        state = "poisoned"
      }
    }
    ospf {
      network_list {
        broadcast_type      = "broadcast"
        p2mp_nbma           = 0
        network_instance_id = 0
      }
      bfd     = 0
      disable = 0
      cost_cfg {
        cost        = 8965
        instance_id = 0
      }
      dead_interval_cfg {
        dead_interval = 40
        instance_id   = 0
      }
      hello_interval_cfg {
        hello_interval = 10
        instance_id    = 0
      }
      mtu_ignore_cfg {
        mtu_ignore  = 0
        instance_id = 0
      }
      neighbor_cfg {
        neighbor               = "2001:db8:3333:4444:5555:6666:7777:8888"
        neig_inst              = 0
        neighbor_cost          = 60182
        neighbor_poll_interval = 2819677845
        neighbor_priority      = 196
      }
      priority_cfg {
        priority    = 1
        instance_id = 0
      }
      retransmit_interval_cfg {
        retransmit_interval = 5
        instance_id         = 0
      }
      transmit_delay_cfg {
        transmit_delay = 1
        instance_id    = 0
      }
    }
  }
  isis {
    authentication {
      send_only_list {
        send_only = 0
        level     = "level-1"
      }
      mode_list {
        mode  = "md5"
        level = "level-1"
      }
      key_chain_list {
        key_chain = "64"
        level     = "level-1"
      }
    }
    bfd_cfg {
      bfd     = 0
      disable = 0
    }
    circuit_type = "level-1-2"
    csnp_interval_list {
      csnp_interval = 10
      level         = "level-1"
    }
    padding = 1
    hello_interval_list {
      hello_interval = 10
      level          = "level-1"
    }
    hello_interval_minimal_list {
      hello_interval_minimal = 0
      level                  = "level-1"
    }
    hello_multiplier_list {
      hello_multiplier = 3
      level            = "level-1"
    }
    lsp_interval = 33
    mesh_group {
      blocked = 0
    }
    metric_list {
      metric = 10
      level  = "level-1"
    }
    network = "broadcast"
    password_list {
      password = "155"
      level    = "level-1"
    }
    priority_list {
      priority = 64
      level    = "level-1"
    }
    retransmit_interval = 5
    wide_metric_list {
      wide_metric = 10
      level       = "level-1"
    }
  }
  l3_vlan_fwd_disable = 0
  lldp {
    enable_cfg {
      rt_enable = 0
      rx        = 0
      tx        = 0
    }
    notification_cfg {
      notification = 0
      notif_enable = 0
    }
    tx_dot1_cfg {
      tx_dot1_tlvs     = 0
      link_aggregation = 0
      vlan             = 0
    }
    tx_tlvs_cfg {
      tx_tlvs             = 0
      exclude             = 0
      management_address  = 0
      port_description    = 0
      system_capabilities = 0
      system_description  = 0
      system_name         = 0
    }
  }
  load_interval = 300
  lw_4o6 {
    outside = 0
    inside  = 0
  }
  mac_learning = "enable"
  map {
    inside        = 0
    outside       = 0
    map_t_inside  = 0
    map_t_outside = 0
  }
  media_type_copper = 0
  monitor_list {
    monitor      = "input"
    mirror_index = 2
    monitor_vlan = 3828
  }
  mtu = mtu
  nptv6 {
    domain_list {
      domain_name = "51"
      bind_type   = "inside"
    }
  }
  packet_capture_template = "107"
  ping_sweep_detection    = "disable"
  port_breakout           = "4x10G"
  port_scan_detection     = "disable"
  remove_vlan_tag         = 0
  sampling_enable {
    counters1 = "all"
  }
  spanning_tree {
    auto_edge  = 1
    admin_edge = 0
    instance_list {
      instance_start = 696
      mstp_path_cost = 198552
    }
    path_cost = 159265
  }
  speed                     = "auto"
  speed_forced_10g          = 0
  speed_forced_40g          = 0
  traffic_distribution_mode = "sip"
  trap_source               = 0
  trunk_group_list {
    trunk_number  = 3250
    type          = "static"
    admin_key     = 11775
    port_priority = 64388
    udld_timeout_cfg {
      slow = 25
    }
    mode     = "active"
    timeout  = "long"
    user_tag = "121"
  }
  update_l2_info = 0
  user_tag       = "29"
  virtual_wire   = 0
  vlan_learning  = "enable"
}