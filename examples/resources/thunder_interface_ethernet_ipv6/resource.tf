provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_ethernet_ipv6" "thunder_interface_ethernet_ipv6" {
  ifnum       = 2
  inside      = 1
  outside     = 1
  ipv6_enable = 1
  ttl_ignore  = 1
  router_adver {
    action            = "disable"
    hop_limit         = 251
    max_interval      = 601
    min_interval      = 203
    default_lifetime  = 1200
    rate_limit        = 10010
    reachable_time    = 5128
    retransmit_timer  = 36734
    adver_mtu_disable = 1
    prefix_list {
      prefix             = "2001:db8:3333:4444:5555:6666:7777:8888/32"
      not_autonomous     = 1
      not_on_link        = 1
      preferred_lifetime = 60480
      valid_lifetime     = 2592100
    }
    managed_config_action = "disable"
    other_config_action   = "disable"
  }
  rip {
    split_horizon_cfg {
      state = "poisoned"
    }
  }
  ospf {
    network_list {
      broadcast_type      = "point-to-point"
      network_instance_id = 0
    }
    bfd     = 1
    disable = 1
    cost_cfg {
      cost        = 8951
      instance_id = 0
    }
    dead_interval_cfg {
      dead_interval = 4033
      instance_id   = 0
    }
    hello_interval_cfg {
      hello_interval = 1025
      instance_id    = 0
    }
    mtu_ignore_cfg {
      mtu_ignore  = 1
      instance_id = 0
    }
    neighbor_cfg {
      neighbor               = "2001:db8:3333:4444:5555:6666:7777:8888"
      neig_inst              = 0
      neighbor_cost          = 6012
      neighbor_poll_interval = 2819645
      neighbor_priority      = 156
    }
    priority_cfg {
      priority    = 123
      instance_id = 0
    }
    retransmit_interval_cfg {
      retransmit_interval = 3236
      instance_id         = 0
    }
    transmit_delay_cfg {
      transmit_delay = 18940
      instance_id    = 0
    }
  }
}