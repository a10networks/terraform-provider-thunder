provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sflow_polling_http" "thunder_sflow_polling_http" {
  toggle = "enable"
}