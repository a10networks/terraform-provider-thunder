provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_http" "thunder_ddos_zone_template_http" {
  http_tmpl_name = "testing"
  mss_timeout {
    mss_percent        = 3
    number_packets     = 4
    mss_timeout_action = "reset"
  }
  disallow_connect_method = 1
  challenge {
    challenge_cookie_name = "test"
  }
  non_http_bypass = 1
  client_source_ip {
    client_source_ip = 1
    http_header_name = "test"
  }
  request_header {
    timeout               = 3
    header_timeout_action = "reset"
  }
  src {
    rate_limit {
      http_post {
        src_post_rate_limit        = 4
        src_post_rate_limit_action = "reset"
      }
      http_request {
        src_request_rate              = 30
        src_request_rate_limit_action = "reset"
      }
    }
  }
  dst {
    rate_limit {
      http_post {
        dst_post_rate_limit_action = "drop"
      }
      http_request {
        dst_request_rate_limit_action = "ignore"
      }
      response_size {
        less_cfg {
          obj_less      = 33
          obj_less_rate = 34
        }

        response_size_action = "reset"
      }
    }
  }
  slow_read {
    min_window_size  = 33
    min_window_count = 3
    slow_read_action = "ignore"
  }
  out_of_order_queue_size    = 3
  out_of_order_queue_timeout = 3
  idle_timeout {
    idle_timeout_value  = 3
    ignore_zero_payload = 1
    idle_timeout_action = "drop"
  }
  user_tag = "test"
  filter_list {
    http_filter_name = "test"
    http_filter_seq  = 3
    http_header_cfg {
      http_filter_header_regex         = "3"
      http_filter_header_inverse_match = 1
    }
    http_referer_cfg {
      referer_contains_cfg {
        http_filter_referer_contains = "test"
      }

    }
    http_agent_cfg {
      agent_contains_cfg {
        http_filter_agent_contains = "test"
      }

    }
    http_uri_cfg {
      uri_equal_cfg {
        http_filter_uri_equals = "test"
      }

      uri_contains_cfg {
        http_filter_uri_contains = "test"
      }

      uri_starts_cfg {
        http_filter_uri_starts_with = "test"
      }

      uri_ends_cfg {
        http_filter_uri_ends_with = "test"
      }

    }
    dst {
      http_filter_rate_limit = 3
    }
    http_filter_action = "drop"
    user_tag           = "test"
  }

  malformed_http {
    malformed_http                       = "check"
    malformed_http_max_line_size         = 3
    malformed_http_max_num_headers       = 6
    malformed_http_max_req_line_size     = 43
    malformed_http_max_header_name_size  = 3
    malformed_http_max_content_length    = 44
    malformed_http_bad_chunk_mon_enabled = 1
    malformed_http_action                = "drop"
  }
}