provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_enable_management_service_ntp" "thunder_enable_management_service_ntp" {

  acl_v4_list {
    acl_id     = 91
    management = 1
    user_tag   = "63"
  }
}