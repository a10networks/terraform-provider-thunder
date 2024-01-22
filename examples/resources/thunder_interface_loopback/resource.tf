provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_loopback" "thunder_interface_loopback" {
  ifnum = 5
  ip {
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
  ipv6 {
    address_list {
      anycast    = 0
      link_local = 0
    }
    ipv6_enable = 1
    ospf {
      bfd     = 1
      disable = 1
      cost_cfg {
        cost        = 45907
        instance_id = 0
      }
      dead_interval_cfg {
        dead_interval = 40
        instance_id   = 0
      }
      hello_interval_cfg {
        hello_interval = 10
        instance_id    = 0
      }
      mtu_ignore_cfg {
        mtu_ignore  = 0
        instance_id = 0
      }
      priority_cfg {
        priority    = 1
        instance_id = 0
      }
      retransmit_interval_cfg {
        retransmit_interval = 5
        instance_id         = 0
      }
      transmit_delay_cfg {
        transmit_delay = 1
        instance_id    = 0
      }
    }
    rip {
      split_horizon_cfg {
        state = "poisoned"
      }
    }
  }
  isis {
    authentication {
      send_only_list {
        send_only = 1
        level     = "level-1"
      }
      mode_list {
        mode  = "md5"
        level = "level-1"
      }
      key_chain_list {
        key_chain = "109"
        level     = "level-1"
      }
    }
    bfd_cfg {
      bfd     = 1
      disable = 1
    }
    circuit_type = "level-1-2"
    mesh_group {
      blocked = 1
    }
  }
  snmp_server {
    trap_source = 0
  }
  user_tag = "52"
}