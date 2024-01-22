provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_entry_src_port_range" "thunder_ddos_dst_entry_src_port_range" {
  dst_entry_name       = "test"
  deny                 = 1
  protocol             = "udp"
  set_counter_base_val = 39
  src_port_range_end   = 28
  src_port_range_start = 24
  user_tag             = "111"
}