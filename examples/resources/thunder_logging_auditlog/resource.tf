provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_logging_auditlog" "thunder_logging_auditlog" {
  audit_facility = "local0"
  host4          = "15"
  port           = 514
}