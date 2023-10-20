provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_client_ssl" "client-ssl" {
  name = "client-ssl-template"
  certificate_list {
    cert = "cert1"
    key  = "cert1"
  }

  client_certificate      = "Require"
  close_notify            = 1
  disable_sslv3           = 1
  auth_username           = "common-name subject-alt-name-email subject-alt-name-othername"
  auth_username_attribute = "authattribute"
  ca_certs {
    ca_cert          = "ca1"
    client_ocsp      = 1
    client_ocsp_srvr = "OCSP1"
  }
  ocsp_stapling       = 1
  ocspst_ca_cert      = "ca1"
  ocspst_ocsp         = 1
  ocspst_srvr         = "OCSP1"
  ocspst_srvr_days    = 29
  ocspst_srvr_timeout = 20

  req_ca_lists {
    client_certificate_request_ca = "ca1"
  }

  dh_type = "2048"
  ec_list {
    ec = "secp384r1"
  }
  ec_list {
    ec = "secp256r1"
  }

  direct_client_server_auth = 1
  early_data                = 1
  no_anti_replay            = 1
  enable_ssli_ftp_alg       = 23
  enable_tls_alert_logging  = 1
  alert_type                = "fatal"
  local_logging             = 1
  ssli_logging              = 1
  sslilogging               = "disable"

  crl_certs {
    crl = "test_crl.crl"
  }


  forward_proxy_ca_cert  = "cert1"
  forward_proxy_ca_key   = "cert1"
  forward_proxy_alt_sign = 1
  fp_alt_cert            = "cert1"
  fp_alt_key             = "cert1"
  fp_alt_chain_cert      = "cert1"
  forward_proxy_trusted_ca_lists {
    forward_proxy_trusted_ca = "ca1"
  }

  forward_proxy_decrypted_dscp          = 61
  forward_proxy_decrypted_dscp_bypass   = 61
  forward_proxy_verify_cert_fail_action = 1
  verify_cert_fail_action               = "bypass"
  forward_proxy_cert_revoke_action      = 1
  cert_revoke_action                    = "bypass"
  forward_proxy_no_shared_cipher_action = 1
  no_shared_cipher_action               = "bypass"
  forward_proxy_esni_action             = 1
  fp_esni_action                        = "bypass"
  forward_proxy_cert_unknown_action     = 1
  cert_unknown_action                   = "bypass"
  forward_proxy_block_message           = "fp_block_message"
  cache_persistence_list_name           = "mySNI.txt"

  fp_cert_ext_crldp          = "test2.com"
  fp_cert_ext_aia_ca_issuers = "test1.com"
  notbefore                  = 1
  notbeforeday               = 30
  notbeforemonth             = 6
  notbeforeyear              = 2020
  notafter                   = 1
  notafterday                = 30
  notaftermonth              = 6
  notafteryear               = 2035

  forward_proxy_ssl_version        = 31
  forward_proxy_ocsp_disable       = 1
  forward_proxy_crl_disable        = 1
  forward_proxy_cert_cache_timeout = 2147483647
  forward_proxy_cert_cache_limit   = 2147483647
  forward_proxy_cert_expiry        = 1
  expire_hours                     = 1
  forward_proxy_enable             = 1

  handshake_logging_enable         = 1
  forward_proxy_selfsign_redir     = 1
  forward_proxy_failsafe_disable   = 1
  forward_proxy_log_disable        = 1
  fp_cert_fetch_autonat            = "auto"
  fp_cert_fetch_autonat_precedence = 1
  forward_proxy_no_sni_action      = "intercept"
  case_insensitive                 = 1

  ad_group_list                       = "mySNI.txt"
  exception_ad_group_list             = "mySNI.txt"
  exception_user_name_list            = "mySNI.txt"
  exception_sni_cl_name               = "mySNI.txt"
  inspect_list_name                   = "mySNI.txt"
  inspect_certificate_subject_cl_name = "mySNI.txt"
  inspect_certificate_issuer_cl_name  = "mySNI.txt"
  inspect_certificate_san_cl_name     = "mySNI.txt"
  contains_list {

    contains = "test1.com"

  }
  ends_with_list {

    ends_with = "test2.com"

  }
  equals_list {

    equals = "test3.com"

  }
  starts_with_list {
    starts_with = "starts.com"
  }

  certificate_subject_starts_with_list {

    certificate_subject_starts = "example.com"

  }
  certificate_issuer_contains_list {

    certificate_issuer_contains = "example1.com"

  }
  bypass_cert_issuer_class_list_name   = "mySNI.txt"
  exception_certificate_issuer_cl_name = "mySNI.txt"


  certificate_issuer_ends_with_list {

    certificate_issuer_ends_with = "example2.com"

  }
  certificate_issuer_equals_list {

    certificate_issuer_equals = "example3.com"

  }
  certificate_issuer_starts_with_list {

    certificate_issuer_starts = "example4.com"

  }
  certificate_san_contains_list {

    certificate_san_contains = "test.com"

  }
  bypass_cert_san_multi_class_list {
    bypass_cert_san_multi_class_list_name = "mySNI.txt"
  }

  exception_certificate_san_cl_name = "mySNI.txt"
  certificate_san_ends_with_list {
    certificate_san_ends_with = "test.com"
  }

  certificate_san_equals_list {
    certificate_san_equals = "equal.com"
  }

  certificate_san_starts_with_list {
    certificate_san_starts = "start.com"
  }
  certificate_san_contains_list {
    certificate_san_contains = "test.com"
  }
  multi_class_list {
    multi_clist_name = "mySNI.txt"
  }
  user_name_list               = "mySNI.txt"
  client_auth_case_insensitive = 1
  client_auth_contains_list {
    client_auth_contains = "example.com"
  }
  client_auth_class_list = "mySNI.txt"

  client_auth_ends_with_list {
    client_auth_ends_with = "ends.com"
  }

  client_auth_equals_list {
    client_auth_equals = "equal.com"
  }

  client_auth_starts_with_list {
    client_auth_starts_with = "starts.com"
  }

  exception_web_category {
    exception_uncategorized                  = 1
    exception_real_estate                    = 1
    exception_computer_and_internet_security = 1
    exception_financial_services             = 1
    exception_business_and_economy           = 1
    exception_computer_and_internet_info     = 1
    exception_auctions                       = 1
    exception_shopping                       = 1
    exception_cult_and_occult                = 1
    exception_travel                         = 1
    exception_drugs                          = 1
    exception_adult_and_pornography          = 1
    exception_home_and_garden                = 1
    exception_military                       = 1
    exception_social_network                 = 1
    exception_dead_sites                     = 1
    exception_stock_advice_and_tools         = 1
    exception_training_and_tools             = 1
    exception_dating                         = 1
    exception_sex_education                  = 1
    exception_religion                       = 1
    exception_entertainment_and_arts         = 1
    exception_personal_sites_and_blogs       = 1
    exception_legal                          = 1
    exception_local_information              = 1
    exception_streaming_media                = 1
    exception_job_search                     = 1
    exception_gambling                       = 1
    exception_translation                    = 1
    exception_reference_and_research         = 1
    exception_shareware_and_freeware         = 1
    exception_peer_to_peer                   = 1
    exception_marijuana                      = 1
    exception_hacking                        = 1
    exception_games                          = 1
    exception_philosophy_and_politics        = 1
    exception_weapons                        = 1
    exception_pay_to_surf                    = 1
    exception_hunting_and_fishing            = 1
    exception_society                        = 1
    exception_educational_institutions       = 1
    exception_online_greeting_cards          = 1
    exception_sports                         = 1
    exception_swimsuits_and_intimate_apparel = 1
    exception_questionable                   = 1
    exception_kids                           = 1
    exception_hate_and_racism                = 1
    exception_personal_storage               = 1
    exception_violence                       = 1
    exception_keyloggers_and_monitoring      = 1
    exception_search_engines                 = 1
    exception_internet_portals               = 1
    exception_web_advertisements             = 1
    exception_cheating                       = 1
    exception_gross                          = 1
    exception_web_based_email                = 1
    exception_malware_sites                  = 1
    exception_phishing_and_other_fraud       = 1
    exception_proxy_avoid_and_anonymizers    = 1
    exception_spyware_and_adware             = 1
    exception_music                          = 1
    exception_government                     = 1
    exception_nudity                         = 1
    exception_news_and_media                 = 1
    exception_illegal                        = 1
    exception_cdns                           = 1
    exception_internet_communications        = 1
    exception_bot_nets                       = 1
    exception_abortion                       = 1
    exception_health_and_medicine            = 1
    exception_spam_urls                      = 1
    exception_parked_domains                 = 1
    exception_alcohol_and_tobacco            = 1
    exception_image_and_video_search         = 1
    exception_fashion_and_beauty             = 1
    exception_recreation_and_hobbies         = 1
    exception_motor_vehicles                 = 1
    exception_web_hosting_sites              = 1
    exception_nudity_artistic                = 1
    exception_illegal_pornography            = 1
  }
  require_web_category = 1
  web_category {
    uncategorized                  = 1
    real_estate                    = 1
    computer_and_internet_security = 1
    financial_services             = 1
    business_and_economy           = 1
    computer_and_internet_info     = 1
    auctions                       = 1
    shopping                       = 1
    cult_and_occult                = 1
    travel                         = 1
    drugs                          = 1
    adult_and_pornography          = 1
    home_and_garden                = 1
    military                       = 1
    social_network                 = 1
    dead_sites                     = 1
    stock_advice_and_tools         = 1
    training_and_tools             = 1
    dating                         = 1
    sex_education                  = 1
    religion                       = 1
    entertainment_and_arts         = 1
    personal_sites_and_blogs       = 1
    legal                          = 1
    local_information              = 1
    streaming_media                = 1
    job_search                     = 1
    gambling                       = 1
    translation                    = 1
    reference_and_research         = 1
    shareware_and_freeware         = 1
    peer_to_peer                   = 1
    marijuana                      = 1
    hacking                        = 1
    games                          = 1
    philosophy_and_politics        = 1
    weapons                        = 1
    pay_to_surf                    = 1
    hunting_and_fishing            = 1
    society                        = 1
    educational_institutions       = 1
    online_greeting_cards          = 1
    sports                         = 1
    swimsuits_and_intimate_apparel = 1
    questionable                   = 1
    kids                           = 1
    hate_and_racism                = 1
    personal_storage               = 1
    violence                       = 1
    keyloggers_and_monitoring      = 1
    search_engines                 = 1
    internet_portals               = 1
    web_advertisements             = 1
    cheating                       = 1
    gross                          = 1
    web_based_email                = 1
    malware_sites                  = 1
    phishing_and_other_fraud       = 1
    proxy_avoid_and_anonymizers    = 1
    spyware_and_adware             = 1
    music                          = 1
    government                     = 1
    nudity                         = 1
    news_and_media                 = 1
    illegal                        = 1
    cdns                           = 1
    internet_communications        = 1
    bot_nets                       = 1
    abortion                       = 1
    health_and_medicine            = 1
    spam_urls                      = 1
    parked_domains                 = 1
    alcohol_and_tobacco            = 1
    image_and_video_search         = 1
    fashion_and_beauty             = 1
    recreation_and_hobbies         = 1
    motor_vehicles                 = 1
    web_hosting_sites              = 1
    nudity_artistic                = 1
    illegal_pornography            = 1
  }
  web_reputation {
    bypass_low_risk = 1
  }

  forward_proxy_cert_not_ready_action    = "bypass"
  forward_proxy_require_sni_cert_matched = "no-match-action-inspect"
  template_cipher                        = "cipher1"
  server_name_auto_map                   = 1
  sni_enable_log                         = 1
  sni_bypass_missing_cert                = 1
  sni_bypass_expired_cert                = 1
  sni_bypass_explicit_list               = "mySNI.txt"
  sni_bypass_enable_log                  = 1
  session_cache_size                     = 10
  session_cache_timeout                  = 30
  session_ticket_disable                 = 1
  renegotiation_disable                  = 1
  sslv2_bypass_service_group             = "sg_https"
  authorization                          = 1
  auth_sg                                = "aam_service_group"
  auth_sg_dn                             = 1
  auth_sg_filter                         = "abc"
  non_ssl_bypass_service_group           = "sg_https"
  non_ssl_bypass_l4session               = 1
  session_ticket_lifetime                = 2147483647
  ssl_false_start_disable                = 1
  user_tag                               = "clientssl"
  version                                = 33
  dgversion                              = 30
}