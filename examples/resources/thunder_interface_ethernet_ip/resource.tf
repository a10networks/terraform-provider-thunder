provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_ethernet_ip" "thunder_interface_ethernet_ip" {
  ifnum = 2
  dhcp  = 0
  address_list {
    ipv4_address = "10.0.11.17"
    ipv4_netmask = "255.255.255.0"
  }
  allow_promiscuous_vip = 1
  cache_spoofing_port   = 1
  helper_address_list {
    helper_address = "10.10.10.10"
  }
  inside                    = 1
  outside                   = 1
  ttl_ignore                = 1
  slb_partition_redirect    = 1
  generate_membership_query = 1
  query_interval            = 125
  max_resp_time             = 100
  client                    = 1
  unnumbered                = 1
  router {
    isis {
      tag = "79"
    }
  }
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
      cost = 42829
      database_filter_cfg {
        database_filter = "all"
        out             = 1
      }
      dead_interval  = 12565
      disable        = "all"
      hello_interval = 25367
      message_digest_cfg {
        message_digest_key = 215
        md5 {
          md5_value = "a13"
        }
      }
      mtu        = 48727
      mtu_ignore = 1
      network {
        broadcast = 1
      }
      priority            = 132
      retransmit_interval = 5122
      transmit_delay      = 1133
    }
  }
}