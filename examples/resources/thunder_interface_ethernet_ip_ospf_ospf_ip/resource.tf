provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_ethernet_ip_ospf_ospf_ip" "thunder_interface_ethernet_ip_ospf_ospf_ip" {

  ifnum              = 2
  authentication     = 1
  authentication_key = "7"
  cost               = 21765
  database_filter    = "all"
  dead_interval      = 40
  hello_interval     = 10
  ip_addr            = "10.0.11.17"
  message_digest_cfg {
    message_digest_key = 151
    md5_value          = "15"
  }
  mtu_ignore          = 1
  out                 = 1
  priority            = 132
  retransmit_interval = 512
  transmit_delay      = 112
}