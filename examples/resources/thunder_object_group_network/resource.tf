provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_object_group_network" "thunder_object_group_network" {

  net_name    = "test"
  description = "16"
  usage       = "acl"
  user_tag    = "77"
}