provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_logging_lsn_quota_exceeded" "thunder_logging_lsn_quota_exceeded" {
  custom1               = 0
  custom2               = 0
  custom3               = 0
  custom4               = 0
  custom5               = 0
  custom6               = 0
  disable_pool_based    = 0
  imei                  = 0
  imsi                  = 0
  ip_based              = 0
  msisdn                = 0
  with_radius_attribute = 0
}