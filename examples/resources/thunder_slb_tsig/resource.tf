provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_tsig" "thunder_slb_tsig" {
  check = "40"
}