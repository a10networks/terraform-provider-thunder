provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_packet" "thunder_debug_packet" {
  all_sctp_ports = 1
  arp            = 1
  detail         = 1
  ethernet       = 2
  icmp           = 1
  icmpv6         = 1
  interface      = 1
  l3_protocol    = 1
  l4_protocol    = 1
  neighbor       = 1
  port_range     = 29290
  sctp           = 1
  timestamp      = 1
  ve             = 1313
}