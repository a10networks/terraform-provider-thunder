provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_harmony_controller_profile_thunder_mgmt_ip" "thunder_harmony_controller_profile_thunder_mgmt_ip" {
  ip_address = "10.10.10.10"
}