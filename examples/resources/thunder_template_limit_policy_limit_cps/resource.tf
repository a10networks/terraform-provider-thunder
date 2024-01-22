provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_limit_policy_limit_cps" "thunder_template_limit_policy_limit_cps" {

  policy_number = 2
  value         = 22
}