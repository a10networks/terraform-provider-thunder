provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_client_ssl" "thunder_slb_template_client_ssl" {

  name                = "temp1"
  auth_username       = "common-name"
  authorization       = 0
  cert_revoke_action  = "bypass"
  cert_unknown_action = "bypass"
  close_notify        = 1
  dh_type             = "1024"
  disable_sslv3       = 1
  early_data          = 1
  ec_list {
    ec = "secp256r1"
  }
  enable_tls_alert_logging = 0
  expire_hours             = 111

  forward_proxy_block_message            = "s912"
  forward_proxy_cert_cache_limit         = 5288
  forward_proxy_cert_cache_timeout       = 3610
  forward_proxy_cert_expiry              = 1
  forward_proxy_cert_not_ready_action    = "bypass"
  forward_proxy_cert_revoke_action       = 1
  forward_proxy_cert_unknown_action      = 1
  forward_proxy_crl_disable              = 1
  forward_proxy_decrypted_dscp           = 32
  forward_proxy_decrypted_dscp_bypass    = 61
  forward_proxy_enable                   = 1
  forward_proxy_esni_action              = 0
  forward_proxy_failsafe_disable         = 1
  forward_proxy_log_disable              = 1
  forward_proxy_no_shared_cipher_action  = 1
  forward_proxy_no_sni_action            = "intercept"
  forward_proxy_ocsp_disable             = 1
  forward_proxy_require_sni_cert_matched = "no-match-action-inspect"
  forward_proxy_selfsign_redir           = 1
  forward_proxy_ssl_version              = 33
  forward_proxy_verify_cert_fail_action  = 1
  handshake_logging_enable               = 1
  ja3_enable                             = 1
  ja3_insert_http_header                 = "127"
  ja3_reject_max_number_per_host         = 238
  ja3_ttl                                = 453
  local_logging                          = 1
  non_ssl_bypass_l4session               = 1
  non_ssl_bypass_service_group           = "sg112"
  renegotiation_disable                  = 1
  server_name_auto_map                   = 1
  session_cache_size                     = 121
  session_cache_timeout                  = 23511
  session_ticket_disable                 = 1
  session_ticket_lifetime                = 235112
  sni_bypass_enable_log                  = 1
  sni_bypass_expired_cert                = 1
  sni_bypass_missing_cert                = 1
  sni_enable_log                         = 1
  ssl_false_start_disable                = 1
  ssli_logging                           = 1
  sslilogging                            = "all"
  sslv2_bypass_service_group             = "sg112"
  user_tag                               = "120"
  version                                = 0
}