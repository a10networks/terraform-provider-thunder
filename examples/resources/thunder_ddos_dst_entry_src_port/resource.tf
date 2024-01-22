provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_entry_src_port" "thunder_ddos_dst_entry_src_port" {

  deny                  = 1
  outbound_src_tracking = "disable"
  port_num              = 22
  protocol              = "dns-udp"
  user_tag              = "75"
  dst_entry_name        = "test"
}