provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_enable_management_service_https" "thunder_enable_management_service_https" {

  acl_v4_list {
    acl_id = 52

    management    = 1
    all_data_intf = 1
    user_tag      = "77"
  }


}