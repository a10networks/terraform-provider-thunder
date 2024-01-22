provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_src_dynamic_entry_overflow_policy" "thunder_ddos_src_dynamic_entry_overflow_policy" {
  app_type_list {
    protocol = "dns"
    user_tag = "117"
  }
  default_address_type = "ip"
  exceed_log_cfg {
    log_enable        = 1
    with_sflow_sample = 1
  }
  l4_type_list {
    protocol = "tcp"
    deny     = 1
    user_tag = "26"
  }
  log_periodic = 1
  user_tag     = "43"
}