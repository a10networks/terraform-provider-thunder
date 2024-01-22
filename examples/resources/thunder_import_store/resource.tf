provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_import_store" "thunder_import_store" {
  create      = 1
  delete      = 0
  name        = "test"
  remote_file = "test"
}