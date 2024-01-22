provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_event_filter" "thunder_ddos_event_filter" {

  black_list {
    black_list_dst = 1
    black_list_src = 1
  }
  drop {
    drop_src        = 1
    drop_dst        = 1
    drop_black_list = 1
  }
  filter_name = "40"
  l4_type_list {
    protocol = "tcp"
    tcp_auth {
      tcp_auth_init = 1
      tcp_auth_pass = 1
      tcp_auth_fail = 1
    }
    retrans_syn_cfg {
      retrans_syn        = 1
      retrans_syn_exceed = 1
    }
    out_of_seq  = 1
    zero_window = 1
    user_tag    = "37"
  }
  user_tag = "65"
  white_list {
    white_list_dst = 1
    white_list_src = 1
  }

}