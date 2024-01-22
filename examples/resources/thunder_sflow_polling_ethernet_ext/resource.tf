provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sflow_polling_ethernet_ext" "thunder_sflow_polling_ethernet_ext" {
  start = 2
}