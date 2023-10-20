provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_monitor" "test_thunder_slb_template_monitor" {
  id2 = 2
  link_enable_cfg {
    enaeth       = 2
    ena_sequence = 3
  }
  monitor_relation = "monitor-or"
  link_up_cfg {
    linkup_ethernet1  = 2
    link_up_sequence1 = 4
  }
}