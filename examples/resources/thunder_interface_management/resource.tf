provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_interface_management" "Interface_Management_Test" {
  lldp {
    notification_cfg {
      notif_enable = 1
      notification = 1
    }
    enable_cfg {
      tx        = 1
      rx        = 1
      rt_enable = 1
    }
    tx_dot1_cfg {
      link_aggregation = 1
      vlan             = 1
      tx_dot1_tlvs     = 1
    }
    tx_tlvs_cfg {
      tx_tlvs             = 1
      port_description    = 1
      exclude             = 1
      management_address  = 1
      system_name         = 1
      system_capabilities = 1
      system_description  = 1
    }
  }
  flow_control = 1
  broadcast_rate_limit {
    rate                    = 50
    bcast_rate_limit_enable = 1
  }
  ip {
    dhcp                       = 0
    ipv4_address               = "10.64.44.129"
    control_apps_use_mgmt_port = 1
    default_gateway            = "10.64.44.1"
    ipv4_netmask               = "255.255.255.0"
  }
  access_list {
    acl_id = 50
  }
  sampling_enable {
    counters1 = "all"
  }
  ipv6 {
    default_ipv6_gateway = "fe80::1"
    inbound              = 1
    address_type         = "link-local"
    ipv6_addr            = "fe80::/64"
    v6_acl_name          = "Ipv6Acl"
    ipv6 {
      default_ipv6_gateway = "1111::1"
      ipv6_addr            = "1111::2/64"
    }

    action = "enable"

  }
}