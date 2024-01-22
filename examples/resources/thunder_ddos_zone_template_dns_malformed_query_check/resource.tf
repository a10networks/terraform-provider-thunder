provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_dns_malformed_query_check" "thunder_ddos_zone_template_dns_malformed_query_check" {
  name                       = "testing"
  dns_malformed_query_action = "drop"
  non_query_opcode_check     = "disable"
  skip_multi_packet_check    = 1
  validation_type            = "basic-header-check"

}