provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_netflow_monitor_resend_template" "thunder_netflow_monitor_resend_template" {

  name    = "a11"
  records = 10011
  timeout = 18023
}