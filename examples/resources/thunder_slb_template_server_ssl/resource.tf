provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_server_ssl" "server-ssl" {
  name                  = "test_server_ssl"
  session_cache_size    = 100
  renegotiation_disable = "1"
  ca_certs {
    ca_cert          = "ca1"
    server_ocsp_srvr = "OCSP1"
  }
  dh_type = "1024"
  ec_list {
    ec = "secp384r1"
  }
  ec_list {
    ec = "secp256r1"
  }
  enable_tls_alert_logging = 1
  alert_type               = "fatal"
  close_notify             = 1
  enable_ssli_ftp_alg      = 443
  early_data               = 1
  certificate {
    cert = "cert1"
    key  = "cert1"
  }
  forward_proxy_enable     = 1
  handshake_logging_enable = 1
  ocsp_stapling            = 1
  server_certificate_error {
    error_type = "email"
  }
  server_name           = "sni"
  session_cache_timeout = 300
  session_ticket_enable = 1
  ssli_logging          = 1
  sslilogging           = "all"
  cipher_template       = "cipher1"
  use_client_sni        = 1
  version               = 33
  dgversion             = 33
  user_tag              = "serverssl"
}