provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates" "thunder_visibility_packet_capture_global_templates" {


  template_list {
    name           = "test"
    capture_config = "test"
    user_tag       = "test"
    trigger_sys_obj_stats_severity {
      error          = 1
      error_alert    = 1
      error_warning  = 1
      error_critical = 1
      drop           = 1
      drop_alert     = 1
      drop_warning   = 1
      drop_critical  = 1
    }
    trigger_sys_obj_stats_change {
      system_ctr_lib_acct {
        trigger_stats_inc {
          total_nodes_free_failed   = 1
          total_nodes_unlink_failed = 1
        }
      }
      system_radius_server {
        trigger_stats_inc {
          radius_request_dropped        = 1
          request_bad_secret_dropped    = 1
          request_no_key_vap_dropped    = 1
          request_malformed_dropped     = 1
          radius_table_full             = 1
          secret_not_configured_dropped = 1
          ha_standby_dropped            = 1
          ipv6_prefix_length_mismatch   = 1
          invalid_key                   = 1
        }
      }
      system_fpga_drop {
        trigger_stats_inc {
          mrx_drop            = 1
          hrx_drop            = 1
          siz_drop            = 1
          fcs_drop            = 1
          land_drop           = 1
          empty_frag_drop     = 1
          mic_frag_drop       = 1
          ipv4_opt_drop       = 1
          ipv4_frag           = 1
          bad_ip_hdr_len      = 1
          bad_ip_flags_drop   = 1
          bad_ip_ttl_drop     = 1
          no_ip_payload_drop  = 1
          oversize_ip_payload = 1
          bad_ip_payload_len  = 1
          bad_ip_frag_offset  = 1
          bad_ip_chksum_drop  = 1
          icmp_pod_drop       = 1
          tcp_bad_urg_offet   = 1
          tcp_short_hdr       = 1
          tcp_bad_ip_len      = 1
          tcp_null_flags      = 1
          tcp_null_scan       = 1
          tcp_fin_sin         = 1
          tcp_xmas_flags      = 1
          tcp_xmas_scan       = 1
          tcp_syn_frag        = 1
          tcp_frag_hdr        = 1
          tcp_bad_chksum      = 1
          udp_short_hdr       = 1
          udp_bad_ip_len      = 1
          udp_kb_frags        = 1
          udp_port_lb         = 1
          udp_bad_chksum      = 1
          runt_ip_hdr         = 1
          runt_tcpudp_hdr     = 1
          tun_mismatch        = 1
        }
      }

      system_dpdk_stats {
        trigger_stats_inc {
          pkt_drop            = 1
          pkt_lnk_down_drop   = 1
          err_pkt_drop        = 1
          rx_err              = 1
          tx_err              = 1
          tx_drop             = 1
          rx_len_err          = 1
          rx_over_err         = 1
          rx_crc_err          = 1
          rx_frame_err        = 1
          rx_no_buff_err      = 1
          rx_miss_err         = 1
          tx_abort_err        = 1
          tx_carrier_err      = 1
          tx_fifo_err         = 1
          tx_hbeat_err        = 1
          tx_windows_err      = 1
          rx_long_len_err     = 1
          rx_short_len_err    = 1
          rx_align_err        = 1
          rx_csum_offload_err = 1
          io_rx_que_drop      = 1
          io_tx_que_drop      = 1
          io_ring_drop        = 1
          w_tx_que_drop       = 1
          w_link_down_drop    = 1
          w_ring_drop         = 1
        }
      }
      aam_authentication_global {
        trigger_stats_inc {
          misses                   = 1
          open_socket_failed       = 1
          connect_failed           = 1
          create_timer_failed      = 1
          get_socket_option_failed = 1
          aflex_authz_fail         = 1
          authn_failure            = 1
          authz_failure            = 1
          dns_resolve_failed       = 1
        }
      }
      aam_rdns {
        trigger_stats_inc {
          request_dropped  = 1
          response_failure = 1
          response_error   = 1
          response_timeout = 1
        }
      }
      aam_auth_server_ldap {
        trigger_stats_inc {
          admin_bind_failure    = 1
          bind_failure          = 1
          search_failure        = 1
          authorize_failure     = 1
          timeout_error         = 1
          other_error           = 1
          request_dropped       = 1
          response_failure      = 1
          response_error        = 1
          response_timeout      = 1
          job_start_error       = 1
          polling_control_error = 1
          ssl_session_failure   = 1
          pw_change_failure     = 1
        }
      }

      aam_auth_server_win {
        trigger_stats_inc {
          kerberos_timeout_error               = 1
          kerberos_other_error                 = 1
          ntlm_authentication_failure          = 1
          ntlm_proto_negotiation_failure       = 1
          ntlm_session_setup_failed            = 1
          kerberos_request_dropped             = 1
          kerberos_response_failure            = 1
          kerberos_response_error              = 1
          kerberos_response_timeout            = 1
          kerberos_job_start_error             = 1
          kerberos_polling_control_error       = 1
          ntlm_prepare_req_failed              = 1
          ntlm_timeout_error                   = 1
          ntlm_other_error                     = 1
          ntlm_request_dropped                 = 1
          ntlm_response_failure                = 1
          ntlm_response_error                  = 1
          ntlm_response_timeout                = 1
          ntlm_job_start_error                 = 1
          ntlm_polling_control_error           = 1
          kerberos_pw_expiry                   = 1
          kerberos_pw_change_failure           = 1
          kerberos_validate_kdc_failure        = 1
          kerberos_generate_kdc_keytab_failure = 1
          kerberos_delete_kdc_keytab_failure   = 1
        }
      }
      aam_auth_account {
        trigger_stats_inc {
          request_dropped  = 1
          response_failure = 1
          response_error   = 1
          response_timeout = 1
          response_other   = 1
        }
      }
      aam_auth_saml_global {
        trigger_stats_inc {
          acs_authz_fail = 1
          acs_error      = 1
        }
        // trigger_stats_rate {
        //   threshold_exceeded_by = 5
        //   duration              = 60
        //   acs_authz_fail        = 1
        //   acs_error             = 1
        // }
      }
      aam_auth_relay_kerberos {
        trigger_stats_inc {
          timeout_error         = 1
          other_error           = 1
          request_dropped       = 1
          response_failure      = 1
          response_error        = 1
          response_timeout      = 1
          job_start_error       = 1
          polling_control_error = 1
        }
        // trigger_stats_rate {
        //   threshold_exceeded_by = 5
        //   duration              = 60
        //   timeout_error         = 1
        //   other_error           = 1
        //   request_dropped       = 1
        //   response_failure      = 1
        //   response_error        = 1
        //   response_timeout      = 1
        //   job_start_error       = 1
        //   polling_control_error = 1
        // }
      }
      aam_auth_captcha {
        trigger_stats_inc {
          request_dropped       = 1
          response_failure      = 1
          response_error        = 1
          response_timeout      = 1
          json_fail             = 1
          attr_fail             = 1
          timeout_error         = 1
          other_error           = 1
          job_start_error       = 1
          polling_control_error = 1
        }
        // trigger_stats_rate {
        //   threshold_exceeded_by = 5
        //   duration              = 60
        //   request_dropped       = 1
        //   response_failure      = 1
        //   response_error        = 1
        //   response_timeout      = 1
        //   json_fail             = 1
        //   attr_fail             = 1
        //   timeout_error         = 1
        //   other_error           = 1
        //   job_start_error       = 1
        //   polling_control_error = 1
        // }
      }
    }
  }
}