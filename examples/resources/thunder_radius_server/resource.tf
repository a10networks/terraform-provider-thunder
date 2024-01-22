provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_radius_server" "thunder_radius_server" {
  default_privilege_read_write = 0
  host {
    ipv4_list {
      ipv4_addr = "10.10.10.10"
      secret {
        secret_value = "45"
        port_cfg {
          auth_port  = 60871
          acct_port  = 19086
          retransmit = 3
          timeout    = 3
        }
      }
    }
    name_list {
      hostname = "test"
      secret {
        secret_value = "45"
        port_cfg {
          auth_port  = 16394
          acct_port  = 29951
          retransmit = 3
          timeout    = 3
        }
      }
    }
  }
}