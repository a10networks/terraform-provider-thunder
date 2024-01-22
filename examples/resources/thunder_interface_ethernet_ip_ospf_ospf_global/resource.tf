provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_ethernet_ip_ospf_ospf_global" "thunder_interface_ethernet_ip_ospf_ospf_global" {
  ifnum = 2
  authentication_cfg {
    authentication = 1
  }
  authentication_key = "7"
  bfd_cfg {
    bfd     = 1
    disable = 1
  }
  cost = 42912
  database_filter_cfg {
    database_filter = "all"
    out             = 1
  }
  dead_interval  = 12245
  disable        = "all"
  hello_interval = 2352
  message_digest_cfg {
    message_digest_key = 214
    md5 {
      md5_value = "a13"
    }
  }
  mtu        = 40795
  mtu_ignore = 1
  network {
    broadcast = 1
  }
  priority            = 132
  retransmit_interval = 5122
  transmit_delay      = 1122
}