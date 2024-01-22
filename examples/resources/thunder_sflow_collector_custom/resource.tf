provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sflow_collector_custom" "thunder_sflow_collector_custom" {
  ipv4_addr = "10.10.10.10"
  user_tag  = "84"
}