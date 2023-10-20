provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_interface_ethernet_ip" "IntIP" {
  ifnum                     = 3
  allow_promiscuous_vip     = 1
  syn_cookie                = 1
  cache_spoofing_port       = 1
  ttl_ignore                = 1
  dhcp                      = 0
  server                    = 0
  inside                    = 1
  outside                   = 0
  client                    = 1
  generate_membership_query = 1
  slb_partition_redirect    = 1
  max_resp_time             = 30
  query_interval            = 10
  helper_address_list {
    helper_address = "6.6.6.6"
  }
  address_list {
    ipv4_address = "2.2.2.2"
    ipv4_netmask = "255.255.0.0"
  }

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
    ospf_ip_list {
      ip_addr            = "33.5.4.4"
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
    ospf_global {
      cost                = 100
      mtu_ignore          = 1
      retransmit_interval = 50
      dead_interval       = 100
      hello_interval      = 10
      priority            = 140
      network {
        broadcast           = 1
        non_broadcast       = 0
        point_to_point      = 0
        point_to_multipoint = 0
        p2mp_nbma           = 0

      }
      mtu            = 1400
      transmit_delay = 70
      disable        = "all"
      authentication_cfg {
        value          = "null"
        authentication = 1
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
  }
}
