provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_enable_management_service_ping_acl_v4" "thunder_enable_management_service_ping_acl_v4" {

  acl_id        = 30
  all_data_intf = 1

  management = 1

  user_tag = "127"

}