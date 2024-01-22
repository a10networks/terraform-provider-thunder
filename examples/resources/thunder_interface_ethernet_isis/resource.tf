provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_ethernet_isis" "thunder_interface_ethernet_isis" {
  ifnum = 2
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
      key_chain = "63"
      level     = "level-1"
    }
  }
  bfd_cfg {
    bfd     = 1
    disable = 1
  }
  circuit_type = "level-1-2"
  hello_interval_list {
    hello_interval = 101
    level          = "level-1"
  }
  hello_multiplier_list {
    hello_multiplier = 8
    level            = "level-1"
  }
  lsp_interval = 3332
  mesh_group {
    blocked = 1
  }
  metric_list {
    metric = 62
    level  = "level-1"
  }
  network = "broadcast"
}