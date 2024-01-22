provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_src_port_template_dns_query_resolution_check" "thunder_ddos_src_port_template_dns_query_resolution_check" {

  big_response_action   = "default"
  big_response_size     = 153
  domain_lockup_action  = "default"
  session_timeout_value = 10
  name                  = "testTemplate"
}