provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_harmony_controller_profile_tunnel" "thunder_harmony_controller_profile_tunnel" {
  action = "enable"
}