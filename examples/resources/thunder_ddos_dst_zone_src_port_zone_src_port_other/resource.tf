provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_src_port_zone_src_port_other" "thunder_ddos_dst_zone_src_port_zone_src_port_other" {
  zone_name            = "testZone"
  deny                 = 1
  port_other           = "other"
  protocol             = "udp"
  set_counter_base_val = 621

}