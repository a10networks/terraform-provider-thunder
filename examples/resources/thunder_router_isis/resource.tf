provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_isis" "thunder_router_isis" {
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