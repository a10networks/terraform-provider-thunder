provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_monitor_method_smtp" "thunder_health_monitor_method_smtp" {

  name          = "tf_test"
  mail_from     = "93"
  rcpt_to       = "116"
  smtp          = 1
  smtp_domain   = "35"
  smtp_port     = 251
  smtp_starttls = 1
}