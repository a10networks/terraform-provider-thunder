provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_src_entry" "thunder_ddos_src_entry" {

  app_type_list {
    protocol = "dns"
    user_tag = "2"
  }
  bypass      = 1
  description = "59"
  exceed_log_cfg {
    log_enable = 1
  }
  ip_addr = "10.10.10.10"
  l4_type_list {
    protocol = "tcp"
    action   = "permit"
    user_tag = "46"
  }
  log_periodic   = 1
  src_entry_name = "25"
  user_tag       = "89"
}