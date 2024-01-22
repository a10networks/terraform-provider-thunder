provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authorization_policy" "thunder_aam_authorization_policy" {
  name = "test"

  user_tag = "69"
}