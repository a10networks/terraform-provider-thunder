provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_techreport_priority_partition" "thunder_techreport_priority_partition" {

  part_name = "test"
}