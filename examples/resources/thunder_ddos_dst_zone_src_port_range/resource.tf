provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_src_port_range" "thunder_ddos_dst_zone_src_port_range" {
  zone_name = "testZone"
  capture_config {
    capture_config_name = "10"
    capture_config_mode = "drop"
  }
  deny                 = 1
  protocol             = "udp"
  src_port_range_end   = 24
  src_port_range_start = 20
  user_tag             = "56"

}