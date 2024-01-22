provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_trunk_ip" "thunder_interface_trunk_ip" {

  ifnum = 11
  address_list {
    ipv4_address = "10.0.11.17"
    ipv4_netmask = "255.255.255.0"
  }
  allow_promiscuous_vip     = 1
  cache_spoofing_port       = 1
  client                    = 1
  dhcp                      = 0
  generate_membership_query = 1
  helper_address_list {
    helper_address = "10.10.10.10"
  }
  max_resp_time = 100
  nat {
    inside  = 1
    outside = 1
  }
  ospf {
    ospf_global {
      authentication_cfg {
        authentication = 1
      }
      authentication_key = "7"
      bfd_cfg {
        bfd     = 1
        disable = 1
      }
      cost = 59864
      database_filter_cfg {
        database_filter = "all"
        out             = 1
      }
      dead_interval  = 40
      disable        = "all"
      hello_interval = 10
      message_digest_cfg {
        message_digest_key = 60
        md5 {
          md5_value = "7"
        }
      }
      mtu        = 45408
      mtu_ignore = 1
      network {
        broadcast = 1
      }
      priority            = 1
      retransmit_interval = 5
      transmit_delay      = 1
    }
  }
  query_interval = 125
  rip {
    authentication {
      key_chain {
        key_chain = "v22"
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
  router {
    isis {
      tag = "124"
    }
  }
  slb_partition_redirect = 1
  stateful_firewall {
    inside      = 0
    outside     = 0
    access_list = 0
  }
  syn_cookie = 1
  ttl_ignore = 1
  unnumbered = 0
}