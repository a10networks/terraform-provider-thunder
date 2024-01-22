provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sflow_collector_ipv6" "thunder_sflow_collector_ipv6" {
  addr = "2001:db8:3333:4444:5555:6666:7777:8888"
  port = 4428
}