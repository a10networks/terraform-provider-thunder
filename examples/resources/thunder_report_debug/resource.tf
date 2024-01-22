provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_report_debug" "thunder_report_debug" {

  sflow {
    parser    = 0
    stats_oid = 7945
  }
}