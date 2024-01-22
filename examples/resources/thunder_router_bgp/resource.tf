provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_bgp" "thunder_router_bgp" {
  as_number = "300"
  neighbor {
    ipv4_neighbor_list {
      neighbor_ipv4    = "10.1.1.104"
      activate         = 1
      nbr_remote_as    = 104
      allowas_in       = 1
      allowas_in_count = 10
      graceful_restart = 1
    }
  }
}