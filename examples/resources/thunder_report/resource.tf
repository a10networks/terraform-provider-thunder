provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_report" "thunder_report" {

  debug {
    sflow {
      parser    = 0
      stats_oid = 8770
    }
  }
}