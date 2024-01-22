provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_vrid" "thunder_vrrp_a_vrid" {
  vrid_val = 3
  floating_ip {
    ip_address_cfg {
      ip_address = "10.10.20.11"
    }

  }
  preempt_mode {
    threshold = 2
  }
  user_tag = "test"
  blade_parameters {
    priority = 3
    tracking_options {
      interface {
        ethernet      = 2
        priority_cost = 22
      }

      bgp {
        bgp_ipv4_address_cfg {
          bgp_ipv4_address = "10.10.10.10"
          priority_cost    = 22
        }

      }
      gateway {
        ipv4_gateway_list {
          ip_address    = "20.20.20.20"
          priority_cost = 22
        }

      }
    }
  }
}