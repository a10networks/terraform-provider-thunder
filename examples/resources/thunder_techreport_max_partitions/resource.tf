provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_techreport_max_partitions" "thunder_techreport_max_partitions" {
  value = 5
}