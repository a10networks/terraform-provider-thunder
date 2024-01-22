provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_logging_lsn_log_suppression" "thunder_logging_lsn_log_suppression" {

  count1 = 50
  time   = 20

}