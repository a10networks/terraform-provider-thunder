provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_dblb" "test_thunder_slb_template_dblb" {
  name           = "testing_dblb"
  server_version = "MySQL"
  user_tag       = "test_user"
}