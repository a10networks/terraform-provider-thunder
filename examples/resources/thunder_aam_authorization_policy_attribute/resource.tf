provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authorization_policy_attribute" "thunder_aam_authorization_policy_attribute" {

  name                = "test"
  attribute_name      = "test"
  a10_ax_auth_uri     = 0
  a10_dynamic_defined = 0
  any                 = 1
  attr_num            = 10
  custom_attr_type    = 0
  integer_type        = 0

}