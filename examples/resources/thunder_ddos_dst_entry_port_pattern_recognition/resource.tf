provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_entry_port_pattern_recognition" "thunder_ddos_dst_entry_port_pattern_recognition" {

  algorithm                 = "heuristic"
  filter_inactive_threshold = 109
  filter_threshold          = 80
  mode                      = "capture-never-expire"
  sensitivity               = "high"
  dst_entry_name            = "test"
  protocol                  = "tcp"
  port_num                  = 22
}