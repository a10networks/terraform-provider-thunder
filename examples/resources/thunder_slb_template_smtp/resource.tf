provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_smtp" "thunder_slb_template_smtp" {
  name          = "test11"
  server_domain = "10.2.4.8"
  client_domain_switching {
    switching_type = "contains"
    match_string   = "18"
    service_group  = "sg180"
  }
  client_starttls_type = "optional"
  command_disable {
    disable_type = "expn"
  }
  error_code_to_client = 1
  lf_to_crlf           = 0
  server_starttls_type = "optional"
  service_ready_msg    = "ESMTP mail service ready"
  user_tag             = "65"
}