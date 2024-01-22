provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_admin_aws_accesskey" "thunder_admin_aws_accesskey" {
  user          = "admin"
  delete        = 0
  import        = 0
  show          = 0
  use_mgmt_port = 0
}