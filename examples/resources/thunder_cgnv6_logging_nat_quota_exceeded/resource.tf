provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_logging_nat_quota_exceeded" "thunder_cgnv6_logging_nat_quota_exceeded" {
  level = "critical"
}