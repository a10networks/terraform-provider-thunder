provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_peer_group" "thunder_vrrp_a_peer_group" {
  peer {
    ip_peer_address_cfg {
      ip_peer_address = "10.10.10.10"
    }
    ipv6_peer_address_cfg {
      ipv6_peer_address = "2001:db8:3333:4444:5555:6666:7777:8888"
    }
  }
}