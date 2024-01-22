provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_harmony_controller_config_replace" "thunder_harmony_controller_config_replace" {
  status = "enable"
}