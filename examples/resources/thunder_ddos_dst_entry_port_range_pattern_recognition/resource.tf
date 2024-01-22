provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_entry_port_range_pattern_recognition" "thunder_ddos_dst_entry_port_range_pattern_recognition" {

  algorithm                 = "heuristic"
  filter_inactive_threshold = 248
  filter_threshold          = 12
  mode                      = "capture-never-expire"
  sensitivity               = "high"
  dst_entry_name            = "test"
  port_range_end            = 52035
  port_range_start          = 575
  protocol                  = "tcp"
}