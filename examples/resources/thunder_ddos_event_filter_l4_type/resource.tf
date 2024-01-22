provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_event_filter_l4_type" "thunder_ddos_event_filter_l4_type" {
  filter_name = "40"
  out_of_seq  = 1
  protocol    = "tcp"
  retrans_syn_cfg {
    retrans_syn        = 1
    retrans_syn_exceed = 1
  }
  tcp_auth {
    tcp_auth_init = 1
    tcp_auth_pass = 1
    tcp_auth_fail = 1
  }
  user_tag    = "111"
  zero_window = 1

}