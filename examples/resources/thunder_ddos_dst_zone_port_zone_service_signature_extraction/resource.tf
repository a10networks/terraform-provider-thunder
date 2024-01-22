provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_port_zone_service_signature_extraction" "thunder_ddos_dst_zone_port_zone_service_signature_extraction" {
  zone_name   = "test"
  port_num    = "20"
  protocol    = "tcp"
  algorithm   = "heuristic"
  manual_mode = 1

}