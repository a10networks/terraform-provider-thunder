provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_vrrp_peer_group" "Vrrp_Peer_Group_Test" {
  peer {
    ip_peer_address_cfg {
      ip_peer_address = "12.13.14.15"
    }
    ipv6_peer_address_cfg {
      ipv6_peer_address = "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
    }
  }
}
