provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_transparent_acl_template" "test_thunder_slb_transparent_acl_template" {
  name = "template_name"
}