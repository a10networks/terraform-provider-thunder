provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_http" "thunder_ddos_template_http" {
  http_tmpl_name = "test"
  action         = "drop"
  mss_cfg {
    mss_timeout    = 1
    mss_percent    = 2
    number_packets = 22
  }
  disallow_connect_method = 1
  challenge_method        = "javascript"
  challenge_cookie_name   = "test"
  challenge_keep_cookie   = 1
  challenge_interval      = 10
  non_http_bypass         = 1
  malformed_http {
    malformed_http_enabled               = 1
    malformed_http_max_line_size         = 3
    malformed_http_max_num_headers       = 4
    malformed_http_max_req_line_size     = 22
    malformed_http_max_header_name_size  = 2
    malformed_http_max_content_length    = 3
    malformed_http_bad_chunk_mon_enabled = 1
  }
  use_hdr_ip_cfg {
    use_hdr_ip_as_source = 1
    l7_hdr_name          = "test"
  }
  request_header {
    timeout = 3
  }
  post_rate_limit = 333
  request_rate_limit {
    uri {
      contains_cfg {
        url_contains      = "test"
        url_contains_rate = 33
      }
    }
  }
  response_rate_limit {
    obj_size {
      less_cfg {
        obj_less      = 2
        obj_less_rate = 3
      }
    }
  }
  slow_read_drop {
    min_window_size  = 2
    min_window_count = 5
  }
  idle_timeout               = 2
  ignore_zero_payload        = 1
  out_of_order_queue_size    = 3
  out_of_order_queue_timeout = 3
  referer_filter {
    ref_filter_blacklist = 1
    referer_equals_cfg {
      referer_equals = "test"
    }
    referer_contains_cfg {
      referer_contains = "test"
    }
    referer_starts_cfg {
      referer_starts_with = "test"
    }
    referer_ends_cfg {
      referer_ends_with = "test"
    }
  }
  agent_filter {
    agent_filter_blacklist = 1
    agent_equals_cfg {
      agent_equals = "test"
    }
    agent_contains_cfg {
      agent_contains = "test"
    }
    agent_starts_cfg {
      agent_starts_with = "test"
    }
    agent_ends_cfg {
      agent_ends_with = "test"
    }
  }
  user_tag = "testing"
  filter_header_list {
    http_filter_header_seq       = 2
    http_filter_header_regex     = "test"
    http_filter_header_unmatched = 1
    http_filter_header_blacklist = 1
    user_tag                     = "test"
  }
}