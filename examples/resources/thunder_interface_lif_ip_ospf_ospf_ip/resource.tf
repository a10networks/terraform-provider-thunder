provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_lif_ip_ospf_ospf_ip" "thunder_interface_lif_ip_ospf_ospf_ip" {

  ifname             = "test"
  authentication     = 1
  authentication_key = "1"
  cost               = 64948
  database_filter    = "all"
  dead_interval      = 40
  hello_interval     = 10
  ip_addr            = "10.10.10.10"
  message_digest_cfg {
    message_digest_key = 7
    md5_value          = "10"
  }
  mtu_ignore          = 1
  out                 = 1
  priority            = 1
  retransmit_interval = 115
  transmit_delay      = 10
}