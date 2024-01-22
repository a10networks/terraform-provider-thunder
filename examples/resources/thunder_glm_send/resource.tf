provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_glm_send" "thunder_glm_send" {
  ha_status       = 1
  license_request = 1
}