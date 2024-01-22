provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_enable_management_service_ping" "thunder_enable_management_service_ping" {

  acl_v4_list {
    acl_id        = 94
    management    = 1
    all_data_intf = 1
    user_tag      = "70"
  }

}