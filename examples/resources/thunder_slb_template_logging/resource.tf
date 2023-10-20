provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_logging" "test_thunder_slb_template_logging" {
  name          = "test_logs"
  format        = "test.com"
  local_logging = 1
  service_group = "sg1"
  pcre_mask     = "2"
  tcp_proxy     = "default"
  auto          = "auto"
  user_tag      = "test"
}