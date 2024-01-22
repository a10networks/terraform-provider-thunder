provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_default_port" "thunder_ddos_dst_default_port" {
  default_address_type = "ip"
  deny                 = 1
  glid                 = "3"
  port_num             = 80
  protocol             = "dns-tcp"
  user_tag             = "test"
}