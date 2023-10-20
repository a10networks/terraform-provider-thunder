provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_virtual_server_port" "Slb_Virtual_Server_Port_Test" {
  name             = thunder_virtual_server.resourceSlbVirtualServerTest.name
  port_number      = 300
  ha_conn_mirror   = 1
  protocol         = "tcp"
  precedence       = 0
  port_translation = 0
  acl_list {
    acl_id              = "50"
    acl_id_src_nat_pool = "SNAT-Pool1"
    acl_id_seq_num      = 20
  }
  use_default_if_no_server                         = 0
  cpu_compute                                      = 1
  template_tcp                                     = "TCPTemp"
  substitute_source_mac                            = 1
  shared_partition_dynamic_service_template        = 0
  shared_partition_connection_reuse_template       = 0
  when_down                                        = 1
  shared_partition_persist_destination_ip_template = 0
  shared_partition_external_service_template       = 0
  persist_type                                     = "src-dst-ip-swap-persist"
  shared_partition_http_policy_template            = 0
  use_rcv_hop_for_resp                             = 1
  ignore_global                                    = 1
  req_fail                                         = 0
  no_dest_nat                                      = 0
  user_tag                                         = "Virtualport"
  sampling_enable {
    counters1 = "all"
  }
  memory_compute                       = 1
  template_policy                      = "Policy"
  reset_on_server_selection_fail       = 1
  ipinip                               = 1
  no_auto_up_on_aflex                  = 1
  rate                                 = 4000
  gslb_enable                          = 0
  service_group                        = "sg1"
  syn_cookie                           = 1
  alternate_port                       = 0
  alternate_port_number                = 30
  rtp_sip_call_id_match                = 1
  serv_sel_fail                        = 1
  action                               = "enable"
  shared_partition_client_ssl_template = 0
  no_logging                           = 1
  shared_partition_fix_template        = 0
  template_persist_source_ip           = "SRCPersist"
  template_virtual_port                = "80"
  conn_limit                           = 30000
  snat_on_vip                          = 1
  shared_partition_dblb_template       = 0
  shared_partition_http_template       = 0
  force_routing_mode                   = 1
  alt_protocol1                        = "http"
  message_switching                    = 1
  use_alternate_port                   = 1
  eth_fwd                              = 1
  extended_stats                       = 1
  skip_rev_hash                        = 1
  clientip_sticky_nat                  = 1
  secs                                 = 10
  shared_partition_imap_pop3_template  = 0
  eth_rev                              = 2
  stats_data_action                    = "stats-data-disable"
  def_selection_if_pref_failed         = "def-selection-if-pref-failed-disable"
}
