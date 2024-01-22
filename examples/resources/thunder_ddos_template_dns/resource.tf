provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_dns" "thunder_ddos_template_dns" {
  name          = "test"
  action        = "drop"
  dns_any_check = 1
  fqdn_label_len_cfg {
    fqdn_label_length = 1
    label_length      = 2
    fqdn_label_suffix = 2
  }

  fqdn_label_count = 2
  nxdomain_cfg {
    dns_nxdomain_rate_limit        = 1
    dns_nxdomain_rate              = 22
    dns_nxdomain_rate_limit_action = "drop"
  }
  symtimeout_cfg {
    sym_timeout       = 1
    sym_timeout_value = 2
  }
  dns_request_rate_limit {
    type {
      ns_cfg {
        ns          = 1
        dns_ns_rate = 22
      }
    }
  }
  domain_group_rate_exceed_action        = "drop"
  domain_group_rate_per_service          = 1
  query_rate_threshold_for_cache_serving = 22
  allow_query_class {
    allow_any_query_class = 1
  }
  allow_record_type {
    allow_aaaa_type = 1
  }
  user_tag = "test"
  malformed_query_check {
    validation_type         = "disable"
    non_query_opcode_check  = "disable"
    skip_multi_packet_check = 1
  }
}