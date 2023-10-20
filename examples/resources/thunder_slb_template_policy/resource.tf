provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_policy" "template-policy" {
  name    = "policy1"
  timeout = 5
  bw_list_id {
    id             = 3
    service_group  = "sghttp1"
    pbslb_logging  = 1
    pbslb_interval = 10
    fail           = 1
  }
  overlap          = 1
  share            = 1
  full_domain_tree = 1
  user_tag         = "policy testing"
  sampling_enable {
    counters1 = "fwd-policy-dns-unresolved"
  }
  class_list {
    name              = "test_class"
    client_ip_l3_dest = 1
    lid_list {
      user_tag = "lidconfig"
      lidnum   = 1
      response_code_rate_limit {
        code_range_start = 250
        code_range_end   = 300
        threshold        = 3
        period           = 90

      }
      dns64 {
        disable          = 1
        exclusive_answer = 0
      }
      conn_limit            = 1048575
      conn_rate_limit       = 30000
      conn_per              = 2000
      request_limit         = 4000
      request_rate_limit    = 8000
      request_per           = 900
      bw_rate_limit         = 500
      bw_per                = 1000
      over_limit_action     = 1
      action_value          = "forward"
      log                   = 1
      interval              = 30
      direct_action         = 1
      direct_service_group  = "sghttp1"
      direct_pbslb_logging  = 1
      direct_pbslb_interval = 30
      direct_fail           = 1
    }
  }

  forward_policy {
    no_client_conn_reuse = 1
    acos_event_log       = 1
    local_logging        = 1
    require_web_category = 1
    filtering {
      ssli_url_filtering = "bypassed-sni-disable"
    }
    san_filtering {
      ssli_url_filtering_san = "enable-san"
    }
    action_list {
      name               = "Default_Deny"
      action1            = "forward-to-internet"
      fake_sg            = "sghttp1"
      fall_back          = "sghttp1"
      fall_back_snat     = "pool1"
      log                = 0
      drop_message       = "request dropped"
      drop_response_code = 300
      sampling_enable {
        counters1 = "all"
      }
      user_tag = "actiondefaultdeny"
    }

    source_list {
      name = "rule1"
      sampling_enable {
        counters1 = "all"
      }
      match_any              = 1
      match_authorize_policy = "p1"
      priority               = 300
      destination {
        any {
          action = "Default_Deny"
          sampling_enable {
            counters1 = "all"
          }
        }
      }
    }
  }
}