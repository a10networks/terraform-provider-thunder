provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_lif_ip_rip" "thunder_interface_lif_ip_rip" {

  ifname = "test"
  authentication {
    key_chain {
      key_chain = "h126"
    }
  }
  receive_cfg {
    receive = 1
    version = "1"
  }
  receive_packet = 1
  send_cfg {
    send    = 1
    version = "1"
  }
  send_packet = 1
  split_horizon_cfg {
    state = "poisoned"
  }
}