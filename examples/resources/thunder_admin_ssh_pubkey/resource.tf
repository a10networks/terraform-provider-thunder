provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_admin_ssh_pubkey" "thunder_admin_ssh_pubkey" {

  user = "admin5"
  list = 1
}