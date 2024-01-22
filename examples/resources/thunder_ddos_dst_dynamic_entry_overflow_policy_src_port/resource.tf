provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_dynamic_entry_overflow_policy_src_port" "thunder_ddos_dst_dynamic_entry_overflow_policy_src_port" {
  default_address_type = "ip"
  deny                 = 1
  glid                 = "3"
  port_num             = 80
  protocol             = "udp"
  user_tag             = "test"
}