provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_admin_access" "thunder_admin_access" {

  user        = "admin4"
  access_type = "axapi,cli"
}