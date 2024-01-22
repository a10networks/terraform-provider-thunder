provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_partition_template" "thunder_partition_template" {

  partition_name      = "test"
  resource_accounting = "test"
}