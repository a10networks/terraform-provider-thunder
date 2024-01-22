provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_http_filter" "thunder_ddos_zone_template_http_filter" {

  dst {
    http_filter_rate_limit = 12221421
  }
  http_agent_cfg {
    agent_equals_cfg {
      http_filter_agent_equals = "31"
    }
    agent_contains_cfg {
      http_filter_agent_contains = "53"
    }
    agent_starts_cfg {
      http_filter_agent_starts_with = "12"
    }
    agent_ends_cfg {
      http_filter_agent_ends_with = "1"
    }
  }
  http_filter_action = "drop"
  http_filter_name   = "40"
  http_filter_seq    = 122
  http_header_cfg {
    http_filter_header_regex         = "791"
    http_filter_header_inverse_match = 1
  }
  http_referer_cfg {
    referer_equals_cfg {
      http_filter_referer_equals = "29"
    }
    referer_contains_cfg {
      http_filter_referer_contains = "62"
    }
    referer_starts_cfg {
      http_filter_referer_starts_with = "56"
    }
    referer_ends_cfg {
      http_filter_referer_ends_with = "20"
    }
  }
  http_uri_cfg {
    uri_equal_cfg {
      http_filter_uri_equals = "1"
    }
    uri_contains_cfg {
      http_filter_uri_contains = "94"
    }
    uri_starts_cfg {
      http_filter_uri_starts_with = "104"
    }
    uri_ends_cfg {
      http_filter_uri_ends_with = "84"
    }
  }
  user_tag       = "74"
  http_tmpl_name = "testing"
}