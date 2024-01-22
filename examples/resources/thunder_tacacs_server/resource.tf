provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_tacacs_server" "thunder_tacacs_server" {
  monitor  = 1
  interval = 111
  host {
    ipv4_list {
      ipv4_addr = "10.10.10.10"
      secret {
        source_eth   = 2
        secret_value = "test"
        port_cfg {
          port           = 22
          timeout        = 3
          monitor        = 1
          username       = "pramod"
          password       = 1
          password_value = "pramod"
        }
      }
    }

    ipv6_list {
      ipv6_addr = "2001:db8:3333:4444:5555:6666:7777:8888"
      secret {
        source_eth   = 2
        secret_value = "test"
        port_cfg {
          port           = 22
          timeout        = 11
          monitor        = 1
          username       = "suraj123"
          password       = 1
          password_value = "suraj"
        }
      }
    }

    tacacs_hostname_list {
      hostname = "test"
      secret {
        source_eth   = 2
        secret_value = "test"
        port_cfg {
          port           = 22
          timeout        = 2
          monitor        = 1
          username       = "suraj"
          password       = 1
          password_value = "suraj"
        }
      }
    }
  }
}