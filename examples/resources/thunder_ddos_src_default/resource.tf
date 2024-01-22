provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_src_default" "thunder_ddos_src_default" {
  default_address_type = "ip"
  disable              = 1
  age                  = 22
  exceed_log_cfg {
    log_enable = 1
  }
  log_periodic            = 1
  max_dynamic_entry_count = 22
  user_tag                = "test"
  l4_type_list {
    protocol = "tcp"
    deny     = 1
    user_tag = "test"
  }

  app_type_list {
    protocol = "dns"
    user_tag = "test"
  }

}