provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_report_debug_sflow" "thunder_report_debug_sflow" {
  parser    = 0
  stats_oid = 3724
}