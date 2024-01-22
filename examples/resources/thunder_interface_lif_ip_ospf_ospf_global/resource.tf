provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_lif_ip_ospf_ospf_global" "thunder_interface_lif_ip_ospf_ospf_global" {

  ifname = "test"
  authentication_cfg {
    authentication = 1
  }
  authentication_key = "a12"
  bfd_cfg {
    bfd     = 1
    disable = 1
  }
  cost = 16663
  database_filter_cfg {
    database_filter = "all"
    out             = 1
  }
  dead_interval  = 40
  disable        = "all"
  hello_interval = 10
  message_digest_cfg {
    message_digest_key = 159
    md5 {
      md5_value = "c12"
    }
  }
  mtu        = 51087
  mtu_ignore = 1
  network {
    broadcast = 1
  }
  priority            = 1
  retransmit_interval = 510
  transmit_delay      = 23
}