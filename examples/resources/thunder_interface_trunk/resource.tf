provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_trunk" "thunder_interface_trunk" {

  ifnum  = 11
  action = "enable"
  bfd {
    authentication {
      key_id   = 195
      method   = "md5"
      password = "13"
    }
    demand = 0
    echo   = 0
    interval_cfg {
      interval   = 838
      min_rx     = 952
      multiplier = 14
    }
    per_member_port {
      local_address = "10.10.10.10"
    }
  }
  ip {
    address_list {
      ipv4_address = "10.0.11.17"
      ipv4_netmask = "255.255.255.0"
    }
    allow_promiscuous_vip     = 1
    cache_spoofing_port       = 1
    client                    = 1
    dhcp                      = 0
    generate_membership_query = 1
    helper_address_list {
      helper_address = "10.10.10.10"
    }
    max_resp_time = 100
    nat {
      inside  = 1
      outside = 1
    }
    ospf {
      ospf_global {
        authentication_cfg {
          authentication = 1
        }
        authentication_key = "8"
        bfd_cfg {
          bfd     = 1
          disable = 1
        }
        cost = 57908
        database_filter_cfg {
          database_filter = "all"
          out             = 1
        }
        dead_interval  = 40
        disable        = "all"
        hello_interval = 10
        message_digest_cfg {
          message_digest_key = 197
          md5 {
            md5_value = "16"
          }
        }
        mtu                 = 12994
        mtu_ignore          = 1
        priority            = 1
        retransmit_interval = 5
        transmit_delay      = 1
      }
    }
    query_interval = 125
    rip {
      authentication {
        key_chain {
          key_chain = "v22"
        }
      }
      send_packet    = 1
      receive_packet = 1
      send_cfg {
        send    = 1
        version = "1"
      }
      receive_cfg {
        receive = 1
        version = "1"
      }
      split_horizon_cfg {
        state = "poisoned"
      }
    }
    router {
      isis {
        tag = "124"
      }
    }
    slb_partition_redirect = 1
    stateful_firewall {
      inside      = 0
      outside     = 0
      access_list = 0
    }
    syn_cookie = 1
    ttl_ignore = 1
    unnumbered = 0
  }
  ipv6 {
    ipv6_enable = 0
    nat {
      inside  = 1
      outside = 1
    }
    ospf {
      network_list {
        broadcast_type      = "point-to-point"
        network_instance_id = 0
      }
      bfd     = 1
      disable = 1
      cost_cfg {
        cost        = 52860
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
        mtu_ignore  = 1
        instance_id = 0
      }
      neighbor_cfg {
        neighbor               = "2001:db8:3333:4444:5555:6666:7777:8888"
        neig_inst              = 0
        neighbor_cost          = 23684
        neighbor_poll_interval = 22871251
        neighbor_priority      = 65
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
    rip {
      split_horizon_cfg {
        state = "poisoned"
      }
    }
    stateful_firewall {
      inside  = 1
      outside = 0
    }
    ttl_ignore = 1
  }
  isis {
    authentication {
      send_only_list {
        send_only = 1
        level     = "level-1"
      }
      mode_list {
        mode  = "md5"
        level = "level-1"
      }
      key_chain_list {
        key_chain = "10"
        level     = "level-1"
      }
    }
    bfd_cfg {
      bfd     = 1
      disable = 1
    }
    circuit_type = "level-1-2"
    hello_interval_list {
      hello_interval = 101
      level          = "level-1"
    }
    hello_multiplier_list {
      hello_multiplier = 8
      level            = "level-1"
    }
    lsp_interval = 3332
    mesh_group {
      blocked = 1
    }
    metric_list {
      metric = 62
      level  = "level-1"
    }
    network = "broadcast"
  }
}