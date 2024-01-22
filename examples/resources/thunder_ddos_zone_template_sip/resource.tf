provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_sip" "thunder_ddos_zone_template_sip" {
  sip_tmpl_name = "test"
  src {
    sip_request_rate_limit {
      src_sip_rate_action = "drop"
      method {
        invite_cfg {
          invite              = 1
          src_sip_invite_rate = 3
        }
        bye_cfg {
          bye              = 1
          src_sip_bye_rate = 3
        }
      }
    }
  }
  dst {
    sip_request_rate_limit {
      dst_sip_rate_action = "ignore"
      method {
        invite_cfg {
          invite              = 1
          dst_sip_invite_rate = 4
        }
        bye_cfg {
          bye              = 1
          dst_sip_bye_rate = 3
        }
      }
    }
  }
  idle_timeout {
    idle_timeout_value  = 3
    ignore_zero_payload = 1
    idle_timeout_action = "reset"
  }
  user_tag = "test"
  malformed_sip {
    malformed_sip_check                   = "enable-check"
    malformed_sip_max_line_size           = 2
    malformed_sip_max_uri_length          = 3
    malformed_sip_max_header_name_length  = 3
    malformed_sip_max_header_value_length = 4
    malformed_sip_call_id_max_length      = 2
    malformed_sip_sdp_max_length          = 3
    malformed_sip_action                  = "reset"
  }
  filter_header_list {
    sip_filter_name       = "test"
    sip_filter_header_seq = 30
    sip_header_cfg {
      sip_filter_header_regex         = "test"
      sip_filter_header_inverse_match = 1
    }
    sip_filter_action = "ignore"
    user_tag          = "test"
  }
}