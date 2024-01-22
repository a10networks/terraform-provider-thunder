provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_src_port_zone_src_port" "thunder_ddos_dst_zone_src_port_zone_src_port" {
  zone_name            = "testZone"
  deny                 = 1
  port_num             = 944
  protocol             = "dns-udp"
  set_counter_base_val = 366

}