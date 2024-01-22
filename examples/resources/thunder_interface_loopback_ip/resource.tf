provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_loopback_ip" "thunder_interface_loopback_ip" {

  ifnum = 4
  address_list {
    ipv4_address = "10.0.11.17"
    ipv4_netmask = "255.255.255.0"
  }
  ospf {
    ospf_global {
      authentication_cfg {
        authentication = 1
      }
      authentication_key = "8"
      bfd_cfg {
        bfd     = 1
        disable = 1
      }
      cost = 57908
      database_filter_cfg {
        database_filter = "all"
        out             = 1
      }
      dead_interval  = 40
      disable        = "all"
      hello_interval = 10
      message_digest_cfg {
        message_digest_key = 197
        md5 {
          md5_value = "16"
        }
      }
      mtu                 = 12994
      mtu_ignore          = 1
      priority            = 1
      retransmit_interval = 5
      transmit_delay      = 1
    }
  }
  rip {
    authentication {
      key_chain {
        key_chain = "g26"
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
}