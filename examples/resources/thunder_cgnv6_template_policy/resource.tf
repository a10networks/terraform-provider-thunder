provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_template_policy" "thunder_cgnv6_template_policy" {
  name     = "test1"
  user_tag = "57"
}