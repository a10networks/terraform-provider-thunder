provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture" "thunder_visibility_packet_capture" {
  capture_config_list {
    name                             = "test"
    disable                          = 1
    concurrent_captures              = 33
    concurrent_captures_age          = 4
    number_of_packets_per_conn       = 22
    packet_length                    = 130
    file_size                        = 33
    file_count                       = 11
    number_of_packets_per_capture    = 20
    number_of_packets_total          = 22
    enable_continuous_global_capture = 1
    disable_auto_merge               = 1
    keep_pcap_files_after_merge      = 1
    user_tag                         = "37"
  }

  global_templates {
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

  object_templates {
    templ_gtp_plcy_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "110"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        drop_vld_gtp_ie_repeat_count_exceed     = 1
        drop_vld_reserved_field_set             = 1
        drop_vld_tunnel_id_flag                 = 1
        drop_vld_invalid_flow_label_v0          = 1
        drop_vld_invalid_teid                   = 1
        drop_vld_out_of_state                   = 1
        drop_vld_mandatory_information_element  = 1
        drop_vld_mandatory_ie_in_grouped_ie     = 1
        drop_vld_out_of_order_ie                = 1
        drop_vld_out_of_state_ie                = 1
        drop_vld_reserved_information_element   = 1
        drop_vld_version_not_supported          = 1
        drop_vld_message_length                 = 1
        drop_vld_cross_layer_correlation        = 1
        drop_vld_country_code_mismatch          = 1
        drop_vld_gtp_u_spoofed_source_address   = 1
        drop_vld_gtp_bearer_count_exceed        = 1
        drop_vld_gtp_v2_wrong_lbi_create_bearer = 1
        drop_vld_v0_reserved_message_drop       = 1
        drop_vld_v1_reserved_message_drop       = 1
        drop_vld_v2_reserved_message_drop       = 1
        drop_vld_invalid_pkt_len_piggyback      = 1
        drop_vld_sanity_failed_piggyback        = 1
        drop_vld_sequence_num_correlation       = 1
        drop_vld_gtpv0_seqnum_buffer_full       = 1
        drop_vld_gtpv1_seqnum_buffer_full       = 1
        drop_vld_gtpv2_seqnum_buffer_full       = 1
        drop_vld_gtp_invalid_imsi_len_drop      = 1
        drop_vld_gtp_invalid_apn_len_drop       = 1
        drop_vld_protocol_flag_unset            = 1
      }
    }
    interface_ethernet_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "test"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        input_errors  = 1
        crc           = 1
        runts         = 1
        giants        = 1
        output_errors = 1
        collisions    = 1
        giants_output = 1
      }
    }
    interface_tunnel_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "test"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        num_rx_err_pkts = 1
        num_tx_err_pkts = 1
      }
    }
    aam_jwt_authorization_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "test"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        jwt_authorize_failure = 1
        jwt_missing_token     = 1
        jwt_missing_claim     = 1
        jwt_token_expired     = 1
        jwt_signature_failure = 1
        jwt_other_error       = 1
      }
    }
    aam_aaa_policy_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "test"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        error = 1
      }
    }
    aam_auth_logon_http_ins_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "test"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        spn_krb_faiure = 1
      }
    }
    aam_auth_server_ldap_inst_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "32"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        admin_bind_failure  = 1
        bind_failure        = 1
        search_failure      = 1
        authorize_failure   = 1
        timeout_error       = 1
        other_error         = 1
        ssl_session_failure = 1
        pw_change_failure   = 1
      }
    }
    aam_auth_server_ocsp_inst_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "101"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        timeout          = 1
        fail             = 1
        stapling_timeout = 1
        stapling_fail    = 1
      }
    }
    aam_auth_server_rad_inst_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "126"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        authen_failure     = 1
        authorize_failure  = 1
        timeout_error      = 1
        other_error        = 1
        accounting_failure = 1
      }
    }
    aam_auth_server_win_inst_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "109"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        krb_timeout_error          = 1
        krb_other_error            = 1
        krb_pw_expiry              = 1
        krb_pw_change_failure      = 1
        ntlm_proto_nego_failure    = 1
        ntlm_session_setup_failure = 1
        ntlm_prepare_req_error     = 1
        ntlm_auth_failure          = 1
        ntlm_timeout_error         = 1
        ntlm_other_error           = 1
        krb_validate_kdc_failure   = 1
      }
    }
    aam_auth_saml_service_prov_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "25"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        acs_authz_fail = 1
        acs_error      = 1
      }
    }
    aam_auth_saml_id_prov_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "38"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        md_fail  = 1
        acs_fail = 1
      }
    }
    aam_auth_service_group_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "23"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        server_selection_fail_reset = 1
      }
    }
    aam_auth_service_group_mem_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "126"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        curr_conn_overflow = 1
      }
    }
    aam_auth_relay_hbase_inst_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "57"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        no_creds     = 1
        bad_req      = 1
        unauth       = 1
        forbidden    = 1
        not_found    = 1
        server_error = 1
        unavailable  = 1
      }
    }
    aam_auth_relay_form_inst_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "100"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        invalid_srv_rsp = 1
        post_fail       = 1
        invalid_cred    = 1
        bad_req         = 1
        not_fnd         = 1
        error           = 1
        other_error     = 1
      }
    }
    aam_auth_relay_ntlm_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "27"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        failure            = 1
        buffer_alloc_fail  = 1
        encoding_fail      = 1
        insert_header_fail = 1
        parse_header_fail  = 1
        internal_error     = 1
      }
    }
    aam_auth_relay_ws_fed_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "24"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        failure = 1
      }
    }
    aam_auth_captcha_inst_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "57"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        parse_fail    = 1
        json_fail     = 1
        attr_fail     = 1
        timeout_error = 1
        other_error   = 1
      }
    }
    slb_templ_cache_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "16"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        nc_req_header         = 1
        nc_res_header         = 1
        rv_failure            = 1
        content_toobig        = 1
        content_toosmall      = 1
        entry_create_failures = 1
        header_save_error     = 1
      }
    }
    slb_port_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "52"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        es_resp_300 = 1
        es_resp_400 = 1
        es_resp_500 = 1
        resp_3xx    = 1
        resp_4xx    = 1
        resp_5xx    = 1
      }
    }
    slb_vport_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "1"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        total_mf_dns_pkts                   = 1
        es_total_failure_actions            = 1
        compression_miss_no_client          = 1
        compression_miss_template_exclusion = 1
        loc_deny                            = 1
        dnsrrl_total_dropped                = 1
        dnsrrl_bad_fqdn                     = 1
      }
    }
    cgnv6_serv_group_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "75"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        server_selection_fail_drop  = 1
        server_selection_fail_reset = 1
      }
    }
    cgnv6_dns64_vs_port_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "16"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        es_total_failure_actions = 1
      }
    }
    cgnv6_map_trans_domain_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "105"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        inbound_addr_port_validation_failed = 1
        inbound_rev_lookup_failed           = 1
        inbound_dest_unreachable            = 1
        outbound_addr_validation_failed     = 1
        outbound_rev_lookup_failed          = 1
        outbound_dest_unreachable           = 1
        packet_mtu_exceeded                 = 1
        interface_not_configured            = 1
      }
    }
    cgnv6_encap_domain_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "31"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        inbound_addr_port_validation_failed = 1
        inbound_rev_lookup_failed           = 1
        inbound_dest_unreachable            = 1
        outbound_addr_validation_failed     = 1
        outbound_rev_lookup_failed          = 1
        outbound_dest_unreachable           = 1
        packet_mtu_exceeded                 = 1
        interface_not_configured            = 1
      }
    }
    netflow_monitor_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "42"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        nat44_records_sent_failure              = 1
        nat64_records_sent_failure              = 1
        dslite_records_sent_failure             = 1
        session_event_nat44_records_sent_failur = 1
        session_event_nat64_records_sent_failur = 1
        session_event_dslite_records_sent_failu = 1
        session_event_fw4_records_sent_failure  = 1
        session_event_fw6_records_sent_failure  = 1
        port_mapping_nat44_records_sent_failure = 1
        port_mapping_nat64_records_sent_failure = 1
        port_mapping_dslite_records_sent_failur = 1
        netflow_v5_records_sent_failure         = 1
        netflow_v5_ext_records_sent_failure     = 1
        port_batching_nat44_records_sent_failur = 1
        port_batching_nat64_records_sent_failur = 1
        port_batching_dslite_records_sent_failu = 1
        port_batching_v2_nat44_records_sent_fai = 1
        port_batching_v2_nat64_records_sent_fai = 1
        port_batching_v2_dslite_records_sent_fa = 1
        custom_session_event_nat44_creation_rec = 1
        custom_session_event_nat64_creation_rec = 1
        custom_session_event_dslite_creation_re = 1
        custom_session_event_nat44_deletion_rec = 1
        custom_session_event_nat64_deletion_rec = 1
        custom_session_event_dslite_deletion_re = 1
        custom_session_event_fw4_creation_recor = 1
        custom_session_event_fw6_creation_recor = 1
        custom_session_event_fw4_deletion_recor = 1
        custom_session_event_fw6_deletion_recor = 1
        custom_deny_reset_event_fw4_records_sen = 1
        custom_deny_reset_event_fw6_records_sen = 1
        custom_port_mapping_nat44_creation_reco = 1
        custom_port_mapping_nat64_creation_reco = 1
        custom_port_mapping_dslite_creation_rec = 1
        custom_port_mapping_nat44_deletion_reco = 1
        custom_port_mapping_nat64_deletion_reco = 1
        custom_port_mapping_dslite_deletion_rec = 1
        custom_port_batching_nat44_creation_rec = 1
        custom_port_batching_nat64_creation_rec = 1
        custom_port_batching_dslite_creation_re = 1
        custom_port_batching_nat44_deletion_rec = 1
        custom_port_batching_nat64_deletion_rec = 1
        custom_port_batching_dslite_deletion_re = 1
        custom_port_batching_v2_nat44_creation_ = 1
        custom_port_batching_v2_nat64_creation_ = 1
        custom_port_batching_v2_dslite_creation = 1
        custom_port_batching_v2_nat44_deletion_ = 1
        custom_port_batching_v2_nat64_deletion_ = 1
        custom_port_batching_v2_dslite_deletion = 1
        custom_gtp_c_tunnel_event_records_sent_ = 1
        custom_gtp_u_tunnel_event_records_sent_ = 1
        custom_gtp_deny_event_records_sent_fail = 1
        custom_gtp_info_event_records_sent_fail = 1
        custom_fw_iddos_entry_created_records_s = 1
        custom_fw_iddos_entry_deleted_records_s = 1
        custom_fw_sesn_limit_exceeded_records_s = 1
        custom_nat_iddos_l3_entry_created_recor = 1
        custom_nat_iddos_l3_entry_deleted_recor = 1
        custom_nat_iddos_l4_entry_created_recor = 1
        custom_nat_iddos_l4_entry_deleted_recor = 1
      }
    }
    rule_set_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "50"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        unmatched_drops = 1
        deny            = 1
        reset           = 1
      }
    }
    fw_server_port_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "1"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        es_resp_400          = 1
        es_resp_500          = 1
        es_resp_invalid_http = 1
      }
    }
    fw_service_group_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "124"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        server_selection_fail_reset = 1
      }
    }
    fw_service_group_mem_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "50"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        curr_conn_overflow = 1
      }
    }
    dns_vport_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "38"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        dnsrrl_total_dropped           = 1
        total_filter_drop              = 1
        total_max_query_len_drop       = 1
        rcode_notimpl_receive          = 1
        rcode_notimpl_response         = 1
        gslb_query_bad                 = 1
        gslb_response_bad              = 1
        total_dns_filter_type_drop     = 1
        total_dns_filter_class_drop    = 1
        dns_filter_type_a_drop         = 1
        dns_filter_type_aaaa_drop      = 1
        dns_filter_type_cname_drop     = 1
        dns_filter_type_mx_drop        = 1
        dns_filter_type_ns_drop        = 1
        dns_filter_type_srv_drop       = 1
        dns_filter_type_ptr_drop       = 1
        dns_filter_type_soa_drop       = 1
        dns_filter_type_txt_drop       = 1
        dns_filter_type_any_drop       = 1
        dns_filter_type_others_drop    = 1
        dns_filter_class_internet_drop = 1
        dns_filter_class_chaos_drop    = 1
        dns_filter_class_hesiod_drop   = 1
        dns_filter_class_none_drop     = 1
        dns_filter_class_any_drop      = 1
        dns_filter_class_others_drop   = 1
        dns_rpz_action_drop            = 1
        dnsrrl_bad_fqdn                = 1
      }
    }
    smtp_vport_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "50"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        no_proxy                      = 1
        parse_req_fail                = 1
        server_select_fail            = 1
        forward_req_fail              = 1
        forward_req_data_fail         = 1
        snat_fail                     = 1
        send_client_service_not_ready = 1
        recv_server_unknow_reply_code = 1
        read_request_line_fail        = 1
        get_all_headers_fail          = 1
        too_many_headers              = 1
        line_too_long                 = 1
        line_extend_fail              = 1
        line_table_extend_fail        = 1
        parse_request_line_fail       = 1
        insert_resonse_line_fail      = 1
        remove_resonse_line_fail      = 1
        parse_resonse_line_fail       = 1
        server_starttls_fail          = 1
      }
    }
    pop3_vport_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "53"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        svrsel_fail           = 1
        no_route              = 1
        snat_fail             = 1
        line_too_long         = 1
        invalid_start_line    = 1
        unsupported_command   = 1
        bad_sequence          = 1
        rsv_persist_conn_fail = 1
        smp_v6_fail           = 1
        smp_v4_fail           = 1
        insert_tuple_fail     = 1
        cl_est_err            = 1
        ser_connecting_err    = 1
        server_response_err   = 1
      }
    }
    imap_vport_tmpl_list {
      name           = "test"
      capture_config = "test"
      user_tag       = "71"
      trigger_stats_severity {
        error          = 1
        error_alert    = 1
        error_warning  = 1
        error_critical = 1
        drop           = 1
        drop_alert     = 1
        drop_warning   = 1
        drop_critical  = 1
      }
      trigger_stats_inc {
        svrsel_fail            = 1
        no_route               = 1
        snat_fail              = 1
        line_too_long          = 1
        invalid_start_line     = 1
        cant_find_pasv         = 1
        smp_create_fail        = 1
        data_server_conn_fail  = 1
        data_send_fail         = 1
        cant_find_epsv         = 1
        auth_unsupported       = 1
        unsupported_pbsz_value = 1
        unsupported_prot_value = 1
        bad_sequence           = 1
        rsv_persist_conn_fail  = 1
        smp_v6_fail            = 1
        smp_v4_fail            = 1
        insert_tuple_fail      = 1
        cl_est_err             = 1
        ser_connecting_err     = 1
        server_response_err    = 1
        cl_request_err         = 1
      }
    }
  }
}