provider "thunder" {
    address  = var.dut9049
    username = var.username
    password = var.password
}
resource "thunder_admin_password" "test"{
    password_in_module = "admin"
}
