provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_file_metrics" "thunder_visibility_file_metrics" {
  action = "enable"
}