provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_admin_cloud_cred" "thunder_admin_cloud_cred" {
  user          = "admin"
  delete        = 0
  import        = 0
  show          = 1
  type          = "aws-cred"
  use_mgmt_port = 1
}