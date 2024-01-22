provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_policy_dns" "thunder_gslb_policy_dns" {
  name                  = "test"
  action                = 1
  active_only           = 1
  active_only_fail_safe = 1
  backup_alias          = 1
  backup_server         = 1
  cache                 = 1
  aging_time            = 452645
  cname_detect          = 1
  delegation            = 1
  dynamic_preference    = 1
  dynamic_weight        = 1
  external_ip           = 1
  external_soa          = 1
  geoloc_action         = 1
  geoloc_alias          = 1
  geoloc_policy         = 1
  hint                  = "addition"
  ip_replace            = 1
  ipv6 {
    dns_ipv6_option       = "mapping"
    dns_ipv6_mapping_type = "addition"
  }
  logging = "none"
  proxy_block_port_range_list {
    proxy_block_range_from = 100
    proxy_block_range_to   = 110
  }
  server               = 1
  selected_only        = 1
  server_addition_mx   = 0
  server_any           = 0
  server_authoritative = 0
  server_auto_ns       = 0
  server_auto_ptr      = 0
  server_caa           = 0
  server_cname         = 0
  server_full_list     = 0
  server_mode_only     = 0
  server_mx            = 0
  server_naptr         = 0
  server_ns            = 1
  server_ns_list       = 0
  server_ptr           = 0
  server_sec           = 0
  server_srv           = 0
  server_txt           = 0
  sticky               = 1
  sticky_aging_time    = 5
  sticky_ipv6_mask     = 128
  sticky_mask          = "/32"
  ttl                  = 10


}