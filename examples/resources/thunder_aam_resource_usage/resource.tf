provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_resource_usage" "thunder_aam_resource_usage" {
  identity_provider_limit = 66
}