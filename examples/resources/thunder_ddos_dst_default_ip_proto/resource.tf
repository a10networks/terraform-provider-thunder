provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_default_ip_proto" "thunder_ddos_dst_default_ip_proto" {
  default_address_type = "ip"
  deny                 = 1
  glid                 = "3"
  port_num             = 94
  user_tag             = "test"
}