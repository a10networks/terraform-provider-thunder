provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_interface_ip_port" "thunder_ddos_dst_interface_ip_port" {
  addr     = "10.10.10.10"
  deny     = 1
  port_num = 186
  protocol = "tcp"
  user_tag = "121"
}