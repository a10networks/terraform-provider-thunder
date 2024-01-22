provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_src_default_l4_type" "thunder_ddos_src_default_l4_type" {
  default_address_type = "ip"
  deny                 = 1
  protocol             = "tcp"
  user_tag             = "101"
}