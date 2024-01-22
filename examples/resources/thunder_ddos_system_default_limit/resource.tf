provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_system_default_limit" "thunder_ddos_system_default_limit" {
  default_bit_rate_limit      = 14833754
  default_conn_limit          = 11862436
  default_conn_rate_limit     = 1355542
  default_frag_pkt_rate_limit = 4844071
  default_over_limit_action {
    drop = 1
  }
  default_pkt_rate_limit = 10636
  limit_type             = "dst-entry"
  user_tag               = "102"
}