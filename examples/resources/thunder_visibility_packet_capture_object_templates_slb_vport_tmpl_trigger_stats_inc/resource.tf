provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_slb_vport_tmpl_trigger_stats_inc" "thunder_visibility_packet_capture_object_templates_slb_vport_tmpl_trigger_stats_inc" {

  name                                = "test"
  compression_miss_no_client          = 1
  compression_miss_template_exclusion = 1
  dnsrrl_bad_fqdn                     = 1
  dnsrrl_total_dropped                = 1
  es_total_failure_actions            = 1
  loc_deny                            = 1
  total_mf_dns_pkts                   = 1
}