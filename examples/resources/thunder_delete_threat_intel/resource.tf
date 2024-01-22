provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_delete_threat_intel" "thunder_delete_threat_intel" {
  database = "webroot"
}