provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_lldp" "thunder_network_lldp" {
  enable_cfg {
    enable = 1
    rx     = 1
    tx     = 1
  }
  management_address {
    dns_list {
      dns = "20"
      interface {
        ethernet = 1
      }
    }
    ipv4_addr_list {
      ipv4 = "10.10.10.10"
      interface_ipv4 {
        ipv4_eth = 1
      }
    }
    ipv6_addr_list {
      ipv6 = "f5fd:1fcc:28e2:a236:a282:a957:fc64:32c7"
      interface_ipv6 {
        ipv6_eth = 1
      }
    }
  }
  notification_cfg {
    notification = 0
    interval     = 30
  }
  system_description = "20"
  system_name        = "46"
  tx_set {
    fast_count    = 4
    fast_interval = 1
    hold          = 4
    tx_interval   = 30
    reinit_delay  = 2
  }
}