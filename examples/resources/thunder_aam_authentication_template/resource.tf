provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_template" "thunder_aam_authentication_template" {
  auth_sess_mode = "cookie-based"
  chain {
    chain_server_priority = 3
    chain_sg_priority     = 3
  }
  cookie_domain_group {
    cookie_dmngrp = 27
  }
  cookie_max_age                 = 604800
  cookie_samesite                = "strict"
  cookie_secure_enable           = 1
  forward_logout_disable         = 0
  local_logging                  = 1
  log                            = "use-partition-level-config"
  logout_idle_timeout            = 300
  logout_url                     = "118"
  max_session_time               = 64941
  modify_content_security_policy = 1
  name                           = "test"
  oauth_authorization_server     = "53"
  oauth_client                   = "48"
  redirect_hostname              = "24"
  saml_sp                        = "19"
  type                           = "saml"
  user_tag                       = "43"
}