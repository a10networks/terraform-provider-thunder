provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_virtual_server" "virtual-server" {
  name                    = "VS"
  ip_address              = "30.30.30.1"
  description             = "virtual-server"
  enable_disable_action   = "disable-when-all-ports-down"
  redistribution_flagged  = 1
  vport_disable_action    = "drop-packet"
  arp_disable             = 1
  template_policy         = "temp_policy"
  template_virtual_server = "vs_temp"
  template_logging        = "nat1"
  template_scaleout       = "temp_scaleout"
  stats_data_action       = "stats-data-enable"
  extended_stats          = 1
  disable_vip_adv         = 1
  ha_dynamic              = 1
  redistribute_route_map  = "routeredistribute"
  user_tag                = "vs"
  port_list {
    port_number                    = 25
    protocol                       = "smtp"
    name                           = "smtp"
    conn_limit                     = 1
    reset                          = 0
    no_logging                     = 1
    action                         = "enable"
    def_selection_if_pref_failed   = "def-selection-if-pref-failed"
    skip_rev_hash                  = 1
    message_switching              = 1
    force_routing_mode             = 1
    rate                           = 1
    secs                           = 1
    reset_on_server_selection_fail = 1
    clientip_sticky_nat            = 1
    extended_stats                 = 1
    snat_on_vip                    = 1
    stats_data_action              = "stats-data-enable"
    syn_cookie                     = 0
    attack_detection               = 1
    acl_list {
      acl_name = "acl1"
    }
    template_policy              = "temp_policy"
    no_auto_up_on_aflex          = 1
    scaleout_bucket_count        = 1
    scaleout_device_group        = 1
    auto                         = 0
    service_group                = "sg_tcp"
    ipinip                       = 0
    ip_map_list                  = "map1"
    rtp_sip_call_id_match        = 1
    use_rcv_hop_for_resp         = 1
    persist_type                 = "src-dst-ip-swap-persist"
    eth_fwd                      = 1
    eth_rev                      = 1
    template_persist_source_ip   = "sip1"
    template_smtp                = "smtp1"
    template_server_ssl          = "serverssl1"
    template_virtual_port        = "vp1"
    template_tcp_proxy           = "tcpproxy1"
    use_default_if_no_server     = 0
    no_dest_nat                  = 0
    cpu_compute                  = 1
    memory_compute               = 1
    substitute_source_mac        = 1
    ignore_global                = 1
    aflex_table_entry_syn_enable = 1
    reply_acme_challenge         = 0
    user_tag                     = "vport25smtp"
    sampling_enable {
      counters1 = "curr_conn"
    }
    packet_capture_template = "temp_vport"

  }
}
