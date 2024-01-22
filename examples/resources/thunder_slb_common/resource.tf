provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_common" "thunder_slb_common" {
  aflex_table_entry_aging_interval = 1
  aflex_table_entry_sync {
    aflex_table_entry_sync_enable        = 0
    aflex_table_entry_sync_max_key_len   = 1000
    aflex_table_entry_sync_max_value_len = 1000
    aflex_table_entry_sync_min_lifetime  = 0
  }
  after_disable              = 0
  allow_in_gateway_mode      = 0
  attack_resp_code           = 410
  auto_nat_no_ip_refresh     = "enable"
  auto_translate_port        = 0
  buff_thresh                = 0
  buff_thresh_hw_buff        = 814147845
  buff_thresh_relieve_thresh = 204157524
  buff_thresh_sys_buff_high  = 596889450
  buff_thresh_sys_buff_low   = 1233878464
  cache_expire_time          = 1
  cancel_stream_loop_limit   = 5
  cert_pinning {
    ttl = 144
    candidate_list_feedback_opt_in {
      enable   = 0
      schedule = 0
      week_day = "Monday"
      daily    = 0
    }
  }
  clientside_ip       = "10.10.10.10"
  clientside_ipv6     = "2001:db8:3333:4444:5555:6666:7777:8888"
  compress_block_size = 33117
  conn_rate_limit {
    src_ip_list {
      disable_ipv6_support = disable_ipv6_support
      protocol             = "tcp"
      limit                = 165808
      limit_period         = "100"
      exceed_action        = 0
      lock_out             = 1722
    }
  }
  custom_message        = "589"
  custom_signal_clist   = "14"
  ddos_pkt_count_thresh = 100
  ddos_pkt_size_thresh  = 64
  ddos_protection {
    ipd_enable_toggle = "disable"
    logging {
      ipd_logging_toggle = "enable"
    }
    packets_per_second {
      ipd_tcp = 200
      ipd_udp = 200
    }
  }
  disable_adaptive_resource_check    = 0
  disable_persist_scoring            = 0
  disable_port_masking               = 0
  disable_server_auto_reselect       = 0
  dns_cache_age                      = 300
  dns_cache_age_min_threshold        = 0
  dns_cache_aging_weight             = 1
  dns_cache_enable                   = 0
  dns_cache_entry_size               = 256
  dns_cache_sync                     = 0
  dns_cache_sync_entry_size          = 256
  dns_cache_sync_ttl_threshold       = 0
  dns_cache_ttl_adjustment_enable    = 0
  dns_negative_cache_enable          = 0
  dns_persistent_cache_enable        = 0
  dns_persistent_cache_hit_threshold = 0
  dns_persistent_cache_ttl_threshold = 0
  dns_response_rate_limiting {
    max_table_entries = 564259
  }
  dns_vip_stateless               = 0
  drop_icmp_to_vip_when_vip_down  = 0
  dsr_health_check_enable         = 0
  ecmp_hash                       = "system-default"
  enable_ddos                     = 0
  enable_fast_path_rerouting      = 0
  enable_l7_req_acct              = 0
  entity                          = "server"
  exclude_destination             = "local"
  extended_stats                  = 0
  fast_path_disable               = 0
  gateway_health_check            = 0
  graceful_shutdown               = 6336
  graceful_shutdown_enable        = 0
  health_check_to_all_vip         = 0
  honor_server_response_ttl       = 0
  http_fast_enable                = 0
  hw_compression                  = 0
  hw_syn_rr                       = 430867
  interval                        = 5
  ipv4_offset                     = 0
  l2l3_trunk_lb_disable           = 0
  log_for_reset_unknown_conn      = 0
  low_latency                     = 0
  max_buff_queued_per_conn        = 1000
  max_http_header_count           = 90
  max_local_rate                  = 32
  max_persistent_cache            = max_persistent_cache
  max_remote_rate                 = 15000
  monitor_mode_enable             = 0
  msl_time                        = 2
  mss_table                       = 536
  multi_cpu                       = 0
  n5_new                          = 0
  n5_old                          = 0
  ngwaf_proxy_ipv4                = "10.10.10.10"
  ngwaf_proxy_port                = 56626
  no_auto_up_on_aflex             = 0
  odd_even_nat_enable             = 0
  one_server_conn_hm_rate         = 15
  override_port                   = 0
  pbslb_entry_age                 = 6
  pbslb_overflow_glid             = "350"
  per_thr_percent                 = 5
  ping_sweep_detection            = "disable"
  pkt_rate_for_reset_unknown_conn = 68960
  player_id_check_enable          = 0
  port                            = 56279
  port_scan_detection             = "disable"
  pre_process_enable              = 0
  qat                             = 0
  quic {
    cid_len          = 4
    signature        = "105"
    signature_len    = 3
    signature_offset = 4
    cpu_offset       = 12
    quic_lb_offset   = 8
    enable_hash      = 0
    enable_signature = 0
  }
  range                               = 3
  range_end                           = 17880
  range_start                         = 25364
  rate_limit_logging                  = 0
  recursive_ns_cache                  = "honor-packet-ttl"
  reset_stale_session                 = 0
  resolve_port_conflict               = 0
  response_type                       = "single-answer"
  scale_out                           = 0
  scale_out_traffic_map               = 0
  serverside_ip                       = "10.10.10.10"
  serverside_ipv6                     = "2001:db8:3333:4444:5555:6666:7777:8888"
  service_group_on_no_dest_nat_vports = "enforce-different"
  show_slb_server_legacy_cmd          = 0
  show_slb_service_group_legacy_cmd   = 0
  show_slb_virtual_server_legacy_cmd  = 0
  snat_gwy_for_l3                     = 0
  snat_on_vip                         = 0
  snat_preserve {
    range {
      port1 = 1025
      port2 = 1025
    }
  }
  sort_res                = 0
  ssl_module_usage_enable = 0
  ssl_n5_delay_tx_enable  = 0
  ssl_ratelimit_cfg {
    disable_rate = 0
  }
  ssli_cert_not_ready_inspect_limit   = 2000
  ssli_cert_not_ready_inspect_timeout = 10
  ssli_silent_termination_enable      = 0
  ssli_sni_hash_enable                = 0
  stateless_sg_multi_binding          = 0
  stats_data_disable                  = 0
  substitute_source_mac               = 0
  timeout                             = 15
  traffic_map_type                    = "vport"
  ttl_threshold                       = 9712904
  use_default_sess_count              = 0
  use_https_proxy                     = 0
  use_mss_tab                         = 0
  vport_global                        = 75
  vport_l3v                           = 125
}