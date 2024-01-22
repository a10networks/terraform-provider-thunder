provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_entry_ip_proto" "thunder_ddos_dst_entry_ip_proto" {

  deny           = 1
  dst_entry_name = "test"

  port_num = 22
  user_tag = "126"

}