provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_trunk_ip_ospf_ospf_global" "thunder_interface_trunk_ip_ospf_ospf_global" {

  ifnum = 11
  authentication_cfg {
    authentication = 1
  }
  authentication_key = "7"
  bfd_cfg {
    bfd     = 1
    disable = 1
  }
  cost = 59864
  database_filter_cfg {
    database_filter = "all"
    out             = 1
  }
  dead_interval  = 40
  disable        = "all"
  hello_interval = 10
  message_digest_cfg {
    message_digest_key = 60
    md5 {
      md5_value = "7"
    }
  }
  mtu        = 45408
  mtu_ignore = 1
  network {
    broadcast = 1
  }
  priority            = 1
  retransmit_interval = 5
  transmit_delay      = 1
}