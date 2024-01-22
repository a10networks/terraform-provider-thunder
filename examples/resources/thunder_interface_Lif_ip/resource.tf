provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_lif_ip" "thunder_interface_lif_ip" {

  ifname = "test"
  address_list {
    ipv4_address = "10.0.11.17"
    ipv4_netmask = "255.255.255.0"
  }
  allow_promiscuous_vip     = 1
  cache_spoofing_port       = 1
  dhcp                      = 0
  generate_membership_query = 1
  inside                    = 1
  max_resp_time             = 100
  ospf {
    ospf_global {
      authentication_cfg {
        authentication = 1
      }
      authentication_key = "a10net2"
      bfd_cfg {
        bfd     = 1
        disable = 1
      }
      cost = 46678
      database_filter_cfg {
        database_filter = "all"
        out             = 1
      }
      dead_interval  = 123
      disable        = "all"
      hello_interval = 10
      message_digest_cfg {
        message_digest_key = 97
        md5 {
          md5_value = "b13"
        }
      }
      mtu        = 47296
      mtu_ignore = 1
      network {
        broadcast = 1
      }
      priority            = 1
      retransmit_interval = 5122
      transmit_delay      = 132
    }
  }
  outside        = 1
  query_interval = 135
  rip {
    authentication {
      key_chain {
        key_chain = "h126"
      }
    }
    send_packet    = 1
    receive_packet = 1
    send_cfg {
      send    = 1
      version = "1"
    }
    receive_cfg {
      receive = 1
      version = "1"
    }
    split_horizon_cfg {
      state = "poisoned"
    }
  }
  unnumbered = 1
}