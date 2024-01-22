provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_loopback_isis" "thunder_interface_loopback_isis" {

  ifnum = 4
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