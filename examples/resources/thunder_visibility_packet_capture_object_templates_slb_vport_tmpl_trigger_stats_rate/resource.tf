provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_slb_vport_tmpl_trigger_stats_rate" "thunder_visibility_packet_capture_object_templates_slb_vport_tmpl_trigger_stats_rate" {

  name                                = "test"
  compression_miss_no_client          = 1
  compression_miss_template_exclusion = 1
  dnsrrl_bad_fqdn                     = 1
  dnsrrl_total_dropped                = 1
  duration                            = 60
  es_total_failure_actions            = 1
  loc_deny                            = 1
  threshold_exceeded_by               = 5
  total_mf_dns_pkts                   = 1
}