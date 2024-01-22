provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_admin_azure_cred" "thunder_admin_azure_cred" {
  user          = "admin"
  delete        = 0
  import        = 0
  show          = 1
  use_mgmt_port = 1
}