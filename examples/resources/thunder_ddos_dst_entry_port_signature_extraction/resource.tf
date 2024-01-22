provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_entry_port_signature_extraction" "thunder_ddos_dst_entry_port_signature_extraction" {
  algorithm      = "heuristic"
  manual_mode    = 1
  port_num       = 22
  protocol       = "tcp"
  dst_entry_name = "test"
}