provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_entry_port_range" "thunder_ddos_dst_entry_port_range" {
  capture_config {
    capture_config_name = "10"
    capture_config_mode = "drop"
  }
  dst_entry_name   = "test"
  protocol         = "tcp"
  deny             = 1
  detection_enable = 1
  enable_top_k     = 1
  port_range_end   = 52035
  port_range_start = 575
  user_tag         = "47"

}