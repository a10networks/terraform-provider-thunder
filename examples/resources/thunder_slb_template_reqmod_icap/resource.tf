provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_reqmod_icap" "test_thunder_slb_template_reqmod_icap" {
  name                      = "test_reqmod"
  allowed_http_methods      = "POST"
  include_protocol_in_uri   = 1
  fail_close                = 1
  min_payload_size          = 2
  preview                   = 3
  service_url               = "test.com"
  service_group             = "sg1"
  tcp_proxy                 = "default"
  disable_http_server_reset = 1
  log_only_allowed_method   = 1
  user_tag                  = "testing"
}