provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_rba_role" "thunder_rba_role" {
  name              = "test1"
  default_privilege = "no-access"
  partition_only    = 0
  user_tag          = "97"
}