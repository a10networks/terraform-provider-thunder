provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_ethernet_ipv6_ospf" "thunder_interface_ethernet_ipv6_ospf" {
  ifnum = 2
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