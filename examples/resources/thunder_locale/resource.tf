provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_locale" "thunder_locale" {
  value = "zh_CN.UTF-8"
}