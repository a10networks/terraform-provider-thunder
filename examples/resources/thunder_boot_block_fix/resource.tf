provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_boot_block_fix" "thunder_boot_block_fix" {
  cf = 0
  hd = 0
}