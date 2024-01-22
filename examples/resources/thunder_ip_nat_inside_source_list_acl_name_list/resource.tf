provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_nat_inside_source_list_acl_name_list" "thunder_ip_nat_inside_source_list_acl_name_list" {

  name = "IPAccessList"
  msl  = 612
  pool = "a24"
}