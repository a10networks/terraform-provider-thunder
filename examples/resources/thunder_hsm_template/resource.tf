provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_hsm_template" "thunder_hsm_template" {
  template_name = "test"
  softhsm_enum  = "softHSM"
}