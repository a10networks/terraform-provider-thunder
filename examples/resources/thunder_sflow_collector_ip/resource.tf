provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sflow_collector_ip" "thunder_sflow_collector_ip" {
  addr = "10.10.10.10"
  port = 38248

}