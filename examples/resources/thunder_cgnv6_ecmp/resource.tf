provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_ecmp" "thunder_cgnv6_ecmp" {
  hashing_type = "4-tuple-hash"
}