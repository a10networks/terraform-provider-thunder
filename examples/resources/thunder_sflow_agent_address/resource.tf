provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sflow_agent_address" "thunder_sflow_agent_address" {
}