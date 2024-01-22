provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_ip_list" "thunder_gslb_ip_list" {
  gslb_ip_list_addr_list {
    id1 = 21
  }
  gslb_ip_list_filename = "test"
  gslb_ip_list_obj_name = "test"
  user_tag              = "test"
}