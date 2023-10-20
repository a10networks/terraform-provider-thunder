provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_respmod_icap" "test_thunder_slb_template_respmod_icap" {
  name                      = "testICAP"
  include_protocol_in_uri   = 1
  fail_close                = 1
  min_payload_size          = 22
  preview                   = 33
  service_url               = "test.com"
  service_group             = "sg1"
  tcp_proxy                 = "default"
  disable_http_server_reset = 1
  log_only_allowed_method   = 1
  user_tag                  = "testing"
}