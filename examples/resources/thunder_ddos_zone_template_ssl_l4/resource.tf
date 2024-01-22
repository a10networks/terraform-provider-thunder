provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_ssl_l4" "thunder_ddos_zone_template_ssl_l4" {
  ssl_l4_tmpl_name = "test"
  renegotiation {
    num_renegotiation   = 3
    ssl_l4_reneg_action = "ignore"
  }
  allow_non_tls = 1
  src {
    rate_limit {
      request {
        src_request_rate_limit        = 3
        src_request_rate_limit_action = "reset"
      }
    }
  }
  dst {
    rate_limit {
      request {
        dst_request_rate_limit        = 33
        dst_request_rate_limit_action = "drop"
      }
    }
  }
  user_tag = "test"
  ssl_traffic_check {
    header_inspection        = 1
    header_action            = "ignore"
    check_resumed_connection = 1
  }
}