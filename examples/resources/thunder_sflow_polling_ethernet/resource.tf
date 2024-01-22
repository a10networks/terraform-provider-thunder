provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sflow_polling_ethernet" "thunder_sflow_polling_ethernet" {
  start = start
}