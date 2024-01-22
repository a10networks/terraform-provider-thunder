provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_harmony_controller_profile" "thunder_harmony_controller_profile" {
  host          = "10.10.1.11"
  port          = 31234
  use_mgmt_port = 1
}