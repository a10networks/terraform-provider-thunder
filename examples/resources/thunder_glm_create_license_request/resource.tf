provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_glm_create_license_request" "thunder_glm_create_license_request" {
  create_license_request = 0
}