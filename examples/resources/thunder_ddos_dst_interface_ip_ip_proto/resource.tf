provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_interface_ip_ip_proto" "thunder_ddos_dst_interface_ip_ip_proto" {

  deny     = 1
  port_num = 122
  user_tag = "64"
  addr     = "10.10.10.10"

}