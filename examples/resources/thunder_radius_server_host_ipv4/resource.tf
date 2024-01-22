provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_radius_server_host_ipv4" "thunder_radius_server_host_ipv4" {
  ipv4_addr = "10.10.10.10"
  secret {
    secret_value = "20"
    port_cfg {
      auth_port  = 60314
      acct_port  = 63899
      retransmit = 3
      timeout    = 3
    }
  }
}