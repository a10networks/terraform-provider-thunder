provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_axdebug_apply_config" "thunder_axdebug_apply_config" {

  config_file = "test6"
}