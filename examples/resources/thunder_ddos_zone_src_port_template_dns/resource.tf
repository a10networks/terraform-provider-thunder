provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_src_port_template_dns" "thunder_ddos_zone_src_port_template_dns" {
  name = "test"
  query_resolution_check {
    session_timeout_value = 15
    domain_lockup_action  = "default"
    big_response_size     = 834
    big_response_action   = "default"
  }
  user_tag = "13"

}