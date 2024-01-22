provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_logging_facility" "thunder_logging_facility" {
  facilityname = "local1"
}