provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_reboot" "Reboot" {

  all = 1

}