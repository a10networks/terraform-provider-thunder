provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_axdebug_save_config" "thunder_axdebug_save_config" {
  config_file = "30"
  default     = 0
}