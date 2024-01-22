provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_virtual_port" "virtual-port" {
  name                       = "testvirtualport"
  user_tag                   = "virtualport1"
  reset_unknown_conn         = 1
  ignore_tcp_msl             = 1
  rate                       = 500
  snat_msl                   = 10
  allow_syn_otherflags       = 1
  aflow                      = 1
  conn_limit                 = 100
  conn_limit_reset           = 1
  conn_limit_no_logging      = 1
  drop_unknown_conn          = 1
  reset_l7_on_failover       = 1
  pkt_rate_type              = "src-ip-port"
  pkt_rate_interval          = "100ms"
  pkt_rate_limit_reset       = 9000
  snat_port_preserve         = 1
  conn_rate_limit            = 50
  rate_interval              = "100ms"
  conn_rate_limit_reset      = 1
  conn_rate_limit_no_logging = 1
  non_syn_initiation         = 1
  dscp                       = 50
  log_options                = "no-logging"
  when_rr_enable             = 1
  allow_vip_to_rport_mapping = 1
}