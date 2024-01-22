provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_tftp" "thunder_tftp" {
  blksize = 2000
}