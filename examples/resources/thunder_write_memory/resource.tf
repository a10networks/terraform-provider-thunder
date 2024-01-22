provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_write_memory" "thunder_write_memory" {
  destination = "primary"
  partition   = "shared"
}