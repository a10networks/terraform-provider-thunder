provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_ethernet_lldp" "thunder_interface_ethernet_lldp" {
  ifnum = 2
  enable_cfg {
    rt_enable = 1
    rx        = 1
    tx        = 1
  }
  notification_cfg {
    notification = 1
    notif_enable = 1
  }
  tx_dot1_cfg {
    tx_dot1_tlvs     = 1
    link_aggregation = 1
    vlan             = 1
  }
  tx_tlvs_cfg {
    tx_tlvs             = 1
    exclude             = 1
    management_address  = 1
    port_description    = 1
    system_capabilities = 1
    system_description  = 1
    system_name         = 1
  }
}