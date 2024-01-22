provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_trunk_ipv6" "thunder_interface_trunk_ipv6" {

  ifnum       = 11
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
  router {
    ripng {
      rip = 1
    }
    ospf {
      area_list {
        area_id_num  = 0
        area_id_addr = "10.10.10.10"
        tag          = "60"
        instance_id  = 0
      }
    }
    isis {
      tag = "77"
    }
  }
  router_adver {
    action           = "disable"
    default_lifetime = 1800
    hop_limit        = 255
    max_interval     = 600
    min_interval     = 200
    rate_limit       = 100000
    reachable_time   = 0
    retransmit_timer = 0
    mtu {
      adver_mtu_disable = 1
    }
    prefix_list {
      prefix             = "2001:db8:3333:4444:5555:6666:7777:8888/24"
      not_autonomous     = 1
      not_on_link        = 1
      preferred_lifetime = 6040
      valid_lifetime     = 25920
    }
    managed_config_action = "disable"
    other_config_action   = "disable"
  }
  stateful_firewall {
    inside  = 1
    outside = 0
  }
  ttl_ignore = 1
}