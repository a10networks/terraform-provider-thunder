provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_dynamic_entry_overflow_policy_ip_proto" "thunder_ddos_dst_dynamic_entry_overflow_policy_ip_proto" {
  default_address_type = "ip"
  port_num             = 32
  deny                 = 1
  glid                 = "3"
  user_tag             = "test"
}