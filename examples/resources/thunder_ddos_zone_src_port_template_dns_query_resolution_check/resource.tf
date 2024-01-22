provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_src_port_template_dns_query_resolution_check" "thunder_ddos_zone_src_port_template_dns_query_resolution_check" {
  name                  = "test"
  big_response_action   = "blacklist-src"
  big_response_size     = 691
  domain_lockup_action  = "blacklist-src"
  session_timeout_value = 29

}