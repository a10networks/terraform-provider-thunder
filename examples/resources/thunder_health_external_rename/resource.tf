provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_external_rename" "thunder_health_external_rename" {
  dst_file = "27"
  src_file = "12"
}