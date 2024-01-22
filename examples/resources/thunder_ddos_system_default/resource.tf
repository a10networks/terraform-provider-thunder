provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_system_default" "thunder_ddos_system_default" {
  limit_list {
    limit_type = "dst-entry"
    default_over_limit_action {
      drop = 1
    }
    default_pkt_rate_limit      = 824
    default_bit_rate_limit      = 247
    default_frag_pkt_rate_limit = 7848
    default_conn_limit          = 10650
    default_conn_rate_limit     = 277
    user_tag                    = "test"
  }
}