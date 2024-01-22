provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_capture_config" "thunder_capture_config" {
  count1    = 0
  enable    = 0
  file_size = 0
  name      = "test"
  user_tag  = "63"
}