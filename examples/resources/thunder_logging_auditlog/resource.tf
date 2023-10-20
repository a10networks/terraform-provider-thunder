provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_logging_auditlog" "logging_auditlog_test_ipv4" {
  host4          = "192.168.92.156"
  audit_facility = "local7"
  port           = 19992
}

resource "thunder_logging_auditlog" "logging_auditlog_test_ipv6" {
  host6          = "fd01::11:2"
  audit_facility = "local3"
  port           = 2345
}

resource "thunder_logging_auditlog" "logging_auditlog_test_partition" {
  partition_name = "test"
  audit_facility = "local5"
  port           = 23456
}