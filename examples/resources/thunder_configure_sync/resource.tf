provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_configure_sync" "thunder_configure_sync" {
  address             = "10.64.3.185"
  all_partitions      = 1
  auto_authentication = 1
  private_key         = "87"
  pwd                 = "a10"
  shared              = 0
  type                = "running"
  usr                 = "admin"
}