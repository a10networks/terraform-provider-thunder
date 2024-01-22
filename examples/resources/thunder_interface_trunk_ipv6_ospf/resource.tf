provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_trunk_ipv6_ospf" "thunder_interface_trunk_ipv6_ospf" {

  ifnum = 11
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