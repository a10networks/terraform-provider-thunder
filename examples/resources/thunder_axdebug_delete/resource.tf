provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_axdebug_delete" "thunder_axdebug_delete" {

  config_file = "test9"
}