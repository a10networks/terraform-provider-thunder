provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_zone_local_zone_cfg" "thunder_zone_local_zone_cfg" {

  name       = "test_zone"
  local_type = 1
}