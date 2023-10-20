provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_smtp" "template-smtp" {
  name     = "templatesmtp1"
  user_tag = "smtp1"
  client_domain_switching {
    switching_type = "starts-with"
    match_string   = "smb"
    service_group  = "sghttp1"
  }
  lf_to_crlf = 1
  command_disable {
    disable_type = "expn"
  }
  error_code_to_client = 1
  server_domain        = "www.example.com"
  service_ready_msg    = "testing smtp template for TF"
  client_starttls_type = "optional"
  server_starttls_type = "optional"
  template {
    logging = "log1"
  }
}