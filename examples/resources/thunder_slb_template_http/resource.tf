provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_http" "template-http" {
  name                           = "http1"
  user_tag                       = "http1"
  keep_client_alive              = 1
  req_hdr_wait_time              = 1
  req_hdr_wait_time_val          = 2
  compression_enable             = 1
  cont_100_wait_for_req_complete = 1
  url_hash_persist               = 1
  url_hash_first                 = 5
  use_server_status              = 1
  template {
    logging = "log1"
  }
  compression_auto_disable_on_high_cpu = 1
  compression_content_type {
    content_type = "html"
  }
  compression_exclude_content_type {
    exclude_content_type = "jpeg"
  }
  compression_exclude_uri {
    exclude_uri = "html"
  }
  compression_keep_accept_encoding        = 1
  compression_keep_accept_encoding_enable = 1
  compression_level                       = 9
  compression_minimum_content_length      = 20
  default_charset                         = "iso-8859-1"
  cookie_format                           = "rfc6265"
  prefix                                  = "host"
  cookie_samesite                         = "none"
  failover_url                            = "google.com"
  frame_limit                             = 65535
  host_switching {
    host_switching_type = "contains"
    host_match_string   = "example.com"
    host_service_group  = "sghttp1"
  }
  client_ip_hdr_replace          = 1
  insert_client_ip_header_name   = "Keep-Alive"
  insert_client_ip               = 1
  insert_client_port             = 1
  insert_client_port_header_name = "Upgrade"
  client_port_hdr_replace        = 1
  log_retry                      = 1
  max_concurrent_streams         = 900
  non_http_bypass                = 1
  bypass_sg                      = "sghttp1"
  persist_on_401                 = 1
  redirect                       = 1
  rd_simple_loc                  = "www.redirect.com"
  rd_resp_code                   = "301"
  redirect_rewrite {
    match_list {
      redirect_match = "rewrite.com"
      rewrite_to     = "rewrite1.com"
    }
    redirect_secure      = 1
    redirect_secure_port = 443
  }
  request_header_erase_list {
    request_header_erase = "Proxy-Authenticate"
  }
  request_header_insert_list {

    request_header_insert      = "[Proxy - Authorization]: [Basic YWxhZGRpbjpvcGVuc2VzYW1l]"
    request_header_insert_type = "insert-if-not-exist"
  }
  request_line_case_insensitive = 1
  request_timeout               = 60
  response_content_replace_list {
    response_content_replace = "Connection"
    response_new_string      = "com"
  }
  response_header_erase_list {
    response_header_erase = "Transfer-Encoding"
  }
  response_header_insert_list {

    response_header_insert      = "[Proxy - Authorization]: [Basic YWxhZGRpbjpvcGVuc2VzYW1l]"
    response_header_insert_type = "insert-if-not-exist"
  }
  retry_on_5xx                 = 1
  retry_on_5xx_val             = 1
  strict_transaction_switch    = 1
  term_11client_hdr_conn_close = 1
}