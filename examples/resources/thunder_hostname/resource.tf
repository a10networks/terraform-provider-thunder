provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_hostname" "test_hostname" {
  value = "testingHost"
}