provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_server_ssl" "thunder_slb_template_server_ssl" {
  name         = "test-ssl"
  close_notify = 1
  dh_type      = "1024"
  server_certificate_error {
    error_type = "ignore"
  }
  ec_list {
    ec = "secp256r1"
  }
  enable_ssli_ftp_alg      = 32772
  enable_tls_alert_logging = 0
  forward_proxy_enable     = 1
  handshake_logging_enable = 1
  ocsp_stapling            = 1
  renegotiation_disable    = 1
  session_cache_size       = 111
  session_cache_timeout    = 3404
  session_ticket_enable    = 0
  ssli_logging             = 1
  sslilogging              = "disable"
  use_client_sni           = 1
  user_tag                 = "66"
  version                  = 33
}