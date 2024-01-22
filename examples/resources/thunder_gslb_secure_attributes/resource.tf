provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_secure_attributes" "thunder_gslb_secure_attributes" {
  action = "enable"
}