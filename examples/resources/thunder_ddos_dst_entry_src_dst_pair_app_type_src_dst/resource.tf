provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_entry_src_dst_pair_app_type_src_dst" "thunder_ddos_dst_entry_src_dst_pair_app_type_src_dst" {

  protocol       = "dns"
  dst_entry_name = "test"
  user_tag       = "78"

}