provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_harmony_controller_profile_re_sync" "thunder_harmony_controller_profile_re_sync" {
  analytics_bus   = 0
  schema_registry = 0
}