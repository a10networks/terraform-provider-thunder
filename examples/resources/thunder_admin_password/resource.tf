provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_admin_password" "admin_password" {
  password_in_module = "a11"

}