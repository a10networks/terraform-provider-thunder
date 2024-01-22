provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_dns" "thunder_ddos_zone_template_dns" {

  name                 = "testing"
  dns_any_check        = 1
  dns_any_check_action = "ignore"
  dns_udp_authentication {
    force_tcp_cfg {
      force_tcp                           = 1
      force_tcp_timeout                   = 3
      force_tcp_ignore_client_source_port = 1
    }
    dns_udp_auth_fail_action = "drop"
  }
  fqdn_label_len_cfg {
    label_length             = 22
    fqdn_label_length_action = "ignore"
  }
  fqdn_label_count_cfg {
    label_count             = 3
    fqdn_label_count_action = "ignore"
  }
  src {
    rate_limit {
      nxdomain {
        dns_nxdomain_rate              = 33
        dns_nxdomain_rate_limit_action = "ignore"
      }
      request {
        type {
          a_cfg {
            a              = 1
            src_dns_a_rate = 3
          }
          aaaa_cfg {
            aaaa              = 1
            src_dns_aaaa_rate = 3
          }
          cname_cfg {
            cname              = 1
            src_dns_cname_rate = 3
          }
          mx_cfg {
            mx              = 1
            src_dns_mx_rate = 3
          }
          ns_cfg {
            ns              = 1
            src_dns_ns_rate = 3
          }
        }
        src_dns_request_rate_limit_action = "drop"
      }
    }
  }
  dst {
    rate_limit {
      fqdn {
        dns_fqdn_rate_limit_action = "ignore"
      }
      domain_group_rate_exceed_action = "drop"
      domain_group_rate_per_service   = 1
      request {
        type {
          a_cfg {
            a          = 1
            dns_a_rate = 2
          }
          aaaa_cfg {
            aaaa          = 1
            dns_aaaa_rate = 2
          }
          cname_cfg {
            cname          = 1
            dns_cname_rate = 2
          }
          mx_cfg {
            mx          = 1
            dns_mx_rate = 2
          }
          ns_cfg {
            ns          = 1
            dns_ns_rate = 2
          }
        }
        dst_dns_request_rate_limit_action = "drop"
      }
    }
  }
  symtimeout_cfg {
    sym_timeout       = 1
    sym_timeout_value = 2
  }
  allow_query_class {
    allow_internet_query_class = 1
    allow_csnet_query_class    = 1
    allow_chaos_query_class    = 1
    allow_hesiod_query_class   = 1
    allow_none_query_class     = 1
    allow_any_query_class      = 1
    allow_query_class_action   = "drop"
  }
  allow_record_type {
    allow_a_type             = 1
    allow_aaaa_type          = 1
    allow_cname_type         = 1
    allow_mx_type            = 1
    allow_ns_type            = 1
    allow_srv_type           = 1
    allow_record_type_action = "drop"
  }

  user_tag = "test"
  malformed_query_check {
    validation_type            = "disable"
    non_query_opcode_check     = "disable"
    skip_multi_packet_check    = 1
    dns_malformed_query_action = "ignore"

  }

}