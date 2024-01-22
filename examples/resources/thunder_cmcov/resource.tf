provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cmcov" "thunder_cmcov" {
  dump   = 0
  export = 0
}