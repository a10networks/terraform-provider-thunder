provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_delete_web_category" "thunder_delete_web_category" {
  database = 1
}