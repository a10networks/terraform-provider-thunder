provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_cgnv6_encap_domain_tmpl_trigger_stats_inc" "thunder_visibility_packet_capture_object_templates_cgnv6_encap_domain_tmpl_trigger_stats_inc" {

  name                                = "test"
  inbound_addr_port_validation_failed = 1
  inbound_dest_unreachable            = 1
  inbound_rev_lookup_failed           = 1
  interface_not_configured            = 1
  outbound_addr_validation_failed     = 1
  outbound_dest_unreachable           = 1
  outbound_rev_lookup_failed          = 1
  packet_mtu_exceeded                 = 1
}