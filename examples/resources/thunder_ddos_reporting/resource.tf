provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_reporting" "thunder_ddos_reporting" {
  toggle = "disable-on-limit-reached"
}