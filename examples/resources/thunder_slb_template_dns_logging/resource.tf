provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_dns_logging" "dns_logging" {
  name                        = "logging_template1"
  dns_logging_protocol        = "tcp"
  dns_logging_request_section = "all"
  dns_logging_type            = "query"
  user_tag                    = "dns_logging_template"
  disable                     = 1
}