provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_trunk_ip_ospf_ospf_ip" "thunder_interface_trunk_ip_ospf_ospf_ip" {

  ifnum              = 11
  authentication     = 1
  authentication_key = "1"
  cost               = 30467
  database_filter    = "all"
  dead_interval      = 40
  hello_interval     = 10
  ip_addr            = "10.10.10.10"
  message_digest_cfg {
    message_digest_key = 72
    md5_value          = "2"
  }
  mtu_ignore          = 1
  out                 = 1
  priority            = 1
  retransmit_interval = 5
  transmit_delay      = 1
}