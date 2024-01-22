provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_fail_over_policy_template" "thunder_vrrp_a_fail_over_policy_template" {
  name = "test"
  interface {
    ethernet = 2
    weight   = 22
  }
  gateway {
    gw_ipv4_address_cfg {
      gw_ipv4_address = "20.20.20.20"
      weight          = 22
    }
  }
  bgp {
    bgp_ipv4_address_cfg {
      bgp_ipv4_address = "10.10.10.10"
      weight           = 22
    }

  }
  trunk_cfg {
    trunk           = 3
    weight          = 2
    per_port_weight = 2
  }

  vlan_cfg {
    vlan    = 3
    timeout = 22
    weight  = 3
  }
  user_tag = "test"
}