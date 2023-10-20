provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_ssli" "test_thunder_slb_template_ssli" {
  name     = "test"
  type     = "xmpp"
  user_tag = "test"
}