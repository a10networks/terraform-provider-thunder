provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_src_port_template_dns" "thunder_ddos_src_port_template_dns" {

  query_resolution_check {
    session_timeout_value = 17
    domain_lockup_action  = "default"
    big_response_size     = 889
    big_response_action   = "default"
  }
  user_tag = "9"
  name     = "testTemplate"
}