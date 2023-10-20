provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

provider "thunder" {
  alias     = "L3V_A"
  address   = var.dut9049
  username  = var.username
  password  = var.password
  partition = var.l3v_1
}

resource "thunder_router_isis" "test00" {
  tag                  = "TeSt"
  adjacency_check      = 1
  ignore_lsp_errors    = 0
  is_type              = "level-2-only"
  lsp_refresh_interval = 800
  max_lsp_lifetime     = 65530
  summary_address_list {
    prefix = "10.1.2.3/32"
    level  = "level-2"
  }
  address_family {
    ipv6 {
      adjacency_check = 1
      distance        = 80
      multi_topology_cfg {
        multi_topology = 0
      }
      summary_prefix_list {
        prefix = "10:20::/64"
        level  = "level-2"
      }
    }
  }
}

resource "thunder_router_isis" "test02" {
  provider = thunder.L3V_A
  tag      = "TeSt2"
}
