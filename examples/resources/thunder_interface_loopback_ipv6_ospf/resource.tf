provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_loopback_ipv6_ospf" "thunder_interface_loopback_ipv6_ospf" {

  ifnum   = 4
  bfd     = 1
  disable = 1
  cost_cfg {
    cost        = 45907
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