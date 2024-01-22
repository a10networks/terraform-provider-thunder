provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_gtp_apn_imsi_list" "thunder_template_gtp_apn_imsi_list" {
  name   = "test"
  action = "deny"
  str_list {
    apn            = "86"
    selection_mode = "mobilestation"
    imsi_selection = "15"
  }
  user_tag = "94"
}