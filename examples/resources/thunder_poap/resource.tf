provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_poap" "thunder_poap" {
  action = "disable"
}