provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_common" "resourceSlbCommonTest" {
  port_scan_detection              = "enable"
  ping_sweep_detection             = "enable"
  extended_stats                   = 1
  stats_data_disable               = 1
  graceful_shutdown_enable         = 1
  graceful_shutdown                = 1
  entity                           = "server"
  after_disable                    = 1
  rate_limit_logging               = 1
  max_local_rate                   = 50
  max_remote_rate                  = 16000
  exclude_destination              = "local"
  dsr_health_check_enable          = 1
  aflex_table_entry_aging_interval = 600
  override_port                    = 1
  health_check_to_all_vip          = 1
  reset_stale_session              = 1
  dns_cache_enable                 = 1
  response_type                    = "single-answer"
  ttl_threshold                    = 100
  dns_cache_age                    = 20000
  compress_block_size              = 23000
  dns_cache_entry_size             = 2048
  honor_server_response_ttl        = 1
  max_buff_queued_per_conn         = 2048
  gateway_health_check             = 1
  pkt_rate_for_reset_unknown_conn  = 10485
  log_for_reset_unknown_conn       = 1
  interval                         = 20
  timeout                          = 50
  msl_time                         = 20
  fast_path_disable                = 1
  l2l3_trunk_lb_disable            = 1
  snat_gwy_for_l3                  = 1
  allow_in_gateway_mode            = 1
  disable_server_auto_reselect     = 1
  enable_l7_req_acct               = 1
  disable_adaptive_resource_check  = 1
  snat_on_vip                      = 1
  low_latency                      = 1
  mss_table                        = 723
  resolve_port_conflict            = 1
  no_auto_up_on_aflex              = 1
  hw_syn_rr                        = 7000
  max_http_header_count            = 200
  scale_out                        = 1
  scale_out_traffic_map            = 1
  sort_res                         = 1
  use_mss_tab                      = 1
  auto_nat_no_ip_refresh           = "enable"
  ddos_protection {
    ipd_enable_toggle = "enable"
    logging {
      ipd_logging_toggle = "enable"
    }
    packets_per_second {
      ipd_tcp = 1
      ipd_udp = 0
    }
  }

  drop_icmp_to_vip_when_vip_down      = 1
  player_id_check_enable              = 1
  stateless_sg_multi_binding          = 1
  ecmp_hash                           = "connection-based"
  service_group_on_no_dest_nat_vports = "allow-same"

  snat_preserve {
    range {
      port1 = 1500
      port2 = 1600
    }
  }
  disable_persist_scoring = 1
  ipv4_offset             = 1
  aflex_table_entry_sync {
    aflex_table_entry_sync_enable        = 1
    aflex_table_entry_sync_max_key_len   = 100
    aflex_table_entry_sync_max_value_len = 100
    aflex_table_entry_sync_min_lifetime  = 5000
    uuid                                 = "strings"
  }
  conn_rate_limit {
    src_ip_list {
      protocol      = "tcp"
      log           = 1
      lock_out      = 1
      limit_period  = "100"
      limit         = 10000
      exceed_action = 1
      shared        = 1
    }
  }
  dns_response_rate_limiting {
    max_table_entries = 14000
  }
}
