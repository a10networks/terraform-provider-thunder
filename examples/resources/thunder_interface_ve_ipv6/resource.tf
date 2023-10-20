provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_interface_ve_ipv6" "veipv6" {
  ifnum       = 300
  outside     = 1
  inside      = 0
  v6_acl_name = "Ipv6Acl"
  inbound     = 1
  address_list {
    address_type = "anycast"
    ipv6_addr    = "2021::1/64"
  }
  ipv6_enable = 1
  rip {
    split_horizon_cfg {
      state = "enable"
    }
  }
  ttl_ignore = 1
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
    isis {
      tag = "ipv6ISIS"
    }
    ospf {
      area_list {
        instance_id  = 50
        area_id_addr = "0.0.0.0"
        area_id_num  = 2
        tag          = "ospfv3"
      }
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
      priority    = 150
      instance_id = 50
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
