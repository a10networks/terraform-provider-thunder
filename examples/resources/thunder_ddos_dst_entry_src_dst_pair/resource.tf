provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_entry_src_dst_pair" "thunder_ddos_dst_entry_src_dst_pair" {

  dst_entry_name = "test"
  default        = 1
  bypass         = 1
  exceed_log_cfg {
    log_enable = 1
  }
  log_periodic = 1

  l4_type_src_dst_list {
    protocol = "tcp"
    deny     = 1

    user_tag = "testuser"
  }

  app_type_src_dst_list {
    protocol = "dns"

    user_tag = "testuser"
  }

}