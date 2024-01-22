provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_http" "thunder_slb_template_http" {
  name                                 = "slb-http"
  client_idle_timeout                  = 0
  client_ip_hdr_replace                = 0
  client_port_hdr_replace              = 0
  compression_auto_disable_on_high_cpu = 86
  compression_br_level                 = 1
  compression_br_sliding_window_size   = 24
  compression_content_type {
    content_type = "21"
  }
  compression_enable = 0
  compression_exclude_content_type {
    exclude_content_type = "4"
  }
  compression_exclude_uri {
    exclude_uri = "18"
  }
  compression_keep_accept_encoding        = 0
  compression_keep_accept_encoding_enable = 0
  compression_level                       = 1
  cont_wait_for_req_complete100           = 0
  default_charset                         = "utf-8"
  disallowed_methods_action               = "drop"
  failover_url                            = "92"
  frame_limit                             = 1001
  http_protocol_check {
    get_and_payload             = "drop"
    h2up_content_length_alias   = "drop"
    h2up_with_host_and_auth     = "drop"
    h2up_with_transfer_encoding = "drop"
    header_filter_rule_list {
      seq_num            = 1
      header_name_value  = "19"
      header_value_value = "57"
      action_value       = "drop"
      user_tag           = "47"
    }
    malformed_h2up_header_value          = "drop"
    malformed_h2up_scheme_value          = "drop"
    multiple_content_length              = "drop"
    multiple_transfer_encoding           = "drop"
    transfer_encoding_and_content_length = "drop"
  }
  http2_client_no_snat            = 0
  insert_client_ip                = 1
  insert_client_ip_header_name    = "60"
  insert_client_port              = 1
  insert_client_port_header_name  = "26"
  keep_client_alive               = 0
  log_retry                       = 0
  max_concurrent_streams          = 50
  non_http_bypass                 = 0
  persist_on_401                  = 0
  redirect                        = 0
  retry_on_5xx                    = 0
  retry_on_5xx_per_req_val        = 3
  retry_on_5xx_val                = 3
  server_support_http2_only       = 0
  server_support_http2_only_value = "auto-detect"
  strict_transaction_switch       = 0
  term_11client_hdr_conn_close    = 0
  url_hash_first                  = 0
  url_hash_offset                 = 0
  url_hash_persist                = 0
  use_server_status               = 0
  user_tag                        = "126"
}