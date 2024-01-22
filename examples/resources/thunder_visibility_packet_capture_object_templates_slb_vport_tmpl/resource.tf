provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_slb_vport_tmpl" "thunder_visibility_packet_capture_object_templates_slb_vport_tmpl" {

  name           = "test"
  capture_config = "test"
  trigger_stats_rate {
    threshold_exceeded_by               = 5
    duration                            = 60
    total_mf_dns_pkts                   = 1
    es_total_failure_actions            = 1
    compression_miss_no_client          = 1
    compression_miss_template_exclusion = 1
    loc_deny                            = 1
    dnsrrl_total_dropped                = 1
    dnsrrl_bad_fqdn                     = 1
  }
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
  user_tag = "60"
}