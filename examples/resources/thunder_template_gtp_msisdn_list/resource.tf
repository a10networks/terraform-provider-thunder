provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_gtp_msisdn_list" "thunder_template_gtp_msisdn_list" {
  name   = "test"
  action = "permit"
  str_list {
    msisdn = "20"
  }
  user_tag = "test_user"
}