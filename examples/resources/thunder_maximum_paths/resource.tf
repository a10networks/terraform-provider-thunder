provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_maximum_paths" "thunder_maximum_paths" {
  path = 2
}