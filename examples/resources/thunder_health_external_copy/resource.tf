provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_external_copy" "thunder_health_external_copy" {
  dst_file = "13"
  src_file = "24"
}