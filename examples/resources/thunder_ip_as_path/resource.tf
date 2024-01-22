provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_as_path" "thunder_ip_as_path" {
  access_list = "100"
  action      = "deny"
  value       = "111"
}