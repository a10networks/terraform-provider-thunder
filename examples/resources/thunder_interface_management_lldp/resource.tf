provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_management_lldp" "thunder_interface_management_lldp" {
  enable_cfg {
    rt_enable = 0
    rx        = 0
    tx        = 0
  }
  notification_cfg {
    notification = 0
    notif_enable = 0
  }
  tx_dot1_cfg {
    tx_dot1_tlvs     = 0
    link_aggregation = 0
    vlan             = 0
  }
  tx_tlvs_cfg {
    tx_tlvs             = 0
    exclude             = 0
    management_address  = 0
    port_description    = 0
    system_capabilities = 0
    system_description  = 0
    system_name         = 0
  }
}