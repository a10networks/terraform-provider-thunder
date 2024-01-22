provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_nat_inside_source_list_acl_id_list" "thunder_ip_nat_inside_source_list_acl_id_list" {

  acl_id = 166
  msl    = 475
  pool   = "a24"
}