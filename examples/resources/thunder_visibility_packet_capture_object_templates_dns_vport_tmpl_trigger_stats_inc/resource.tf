provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_dns_vport_tmpl_trigger_stats_inc" "thunder_visibility_packet_capture_object_templates_dns_vport_tmpl_trigger_stats_inc" {

  name                           = "test"
  dns_filter_class_any_drop      = 1
  dns_filter_class_chaos_drop    = 1
  dns_filter_class_hesiod_drop   = 1
  dns_filter_class_internet_drop = 1
  dns_filter_class_none_drop     = 1
  dns_filter_class_others_drop   = 1
  dns_filter_type_a_drop         = 1
  dns_filter_type_aaaa_drop      = 1
  dns_filter_type_any_drop       = 1
  dns_filter_type_cname_drop     = 1
  dns_filter_type_mx_drop        = 1
  dns_filter_type_ns_drop        = 1
  dns_filter_type_others_drop    = 1
  dns_filter_type_ptr_drop       = 1
  dns_filter_type_soa_drop       = 1
  dns_filter_type_srv_drop       = 1
  dns_filter_type_txt_drop       = 1
  dns_rpz_action_drop            = 1
  dnsrrl_bad_fqdn                = 1
  dnsrrl_total_dropped           = 1
  gslb_query_bad                 = 1
  gslb_response_bad              = 1
  rcode_notimpl_receive          = 1
  rcode_notimpl_response         = 1
  total_dns_filter_class_drop    = 1
  total_dns_filter_type_drop     = 1
  total_filter_drop              = 1
  total_max_query_len_drop       = 1
}