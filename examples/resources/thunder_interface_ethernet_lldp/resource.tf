provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_interface_ethernet_lldp" "lldp" {
  ifnum = 5
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