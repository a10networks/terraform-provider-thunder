provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_license_manager_reminder" "thunder_license_manager_reminder" {
  reminder_value = 700293
}