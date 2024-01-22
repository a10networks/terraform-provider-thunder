provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_loopback_ip_ospf_ospf_ip" "thunder_interface_loopback_ip_ospf_ospf_ip" {

  ifnum              = 4
  authentication     = 1
  authentication_key = "5"
  cost               = 42504
  database_filter    = "all"
  dead_interval      = 40
  hello_interval     = 10
  ip_addr            = "10.10.10.10"
  message_digest_cfg {
    message_digest_key = 137
    md5_value          = "8"
  }
  mtu_ignore          = 1
  out                 = 1
  priority            = 1
  retransmit_interval = 5
  transmit_delay      = 1
}