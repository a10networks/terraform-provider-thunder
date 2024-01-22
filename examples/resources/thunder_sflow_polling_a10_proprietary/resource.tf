provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sflow_polling_a10_proprietary" "thunder_sflow_polling_a10_proprietary" {
  export_deprecated_counters = 0
}