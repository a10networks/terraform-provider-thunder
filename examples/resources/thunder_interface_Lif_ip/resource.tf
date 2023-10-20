provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_interface_lif_ip" "lifip" {
  ifname = "lif1"
  address_list {
    ipv4_address = "56.56.56.56"
    ipv4_netmask = "255.255.255.0"
  }
  dhcp                      = 0
  allow_promiscuous_vip     = 1
  inside                    = 1
  outside                   = 0
  generate_membership_query = 1
  query_interval            = 10
  max_resp_time             = 20
  cache_spoofing_port       = 1
  router {
    isis {
      tag = "isis"
    }
  }
  rip {
    authentication {

      mode {
        mode = "md5"
      }
      key_chain {
        key_chain = "ripat"
      }
    }
    send_packet    = 1
    receive_packet = 1
    receive_cfg {
      receive = 1
      version = 2
    }
    send_cfg {
      send    = 1
      version = 2
    }
    split_horizon_cfg {
      state = "enable"
    }
  }
  ospf {
    ospf_global {
      cost                = 120
      dead_interval       = 40
      hello_interval      = 10
      disable             = "all"
      mtu                 = 1400
      mtu_ignore          = 0
      priority            = 200
      retransmit_interval = 40
      transmit_delay      = 200
      authentication_key  = "ospfKey"
      network {
        broadcast           = 1
        non_broadcast       = 0
        p2mp_nbma           = 0
        point_to_multipoint = 0
        point_to_point      = 0
      }
      authentication_cfg {
        authentication = 1
        value          = "null"
      }
      database_filter_cfg {
        database_filter = "all"
        out             = 1
      }
      bfd_cfg {
        bfd     = 1
        disable = 0
      }

    }
    ospf_ip_list {
      ip_addr            = "33.4.4.4"
      authentication     = 1
      value              = "null"
      authentication_key = "oppsf"
      cost               = 500
      database_filter    = "all"
      out                = 1
      dead_interval      = 40
      hello_interval     = 10
      message_digest_cfg {
        message_digest_key = 4
        md5_value          = "md5"

      }
      mtu_ignore          = 0
      priority            = 120
      retransmit_interval = 50
      transmit_delay      = 100
    }
  }
}
