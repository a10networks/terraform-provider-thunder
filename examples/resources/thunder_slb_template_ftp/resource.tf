provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_ftp" "test_thunder_slb_template_ftp" {
  name                 = "test_ftp"
  active_mode_port     = 1
  active_mode_port_val = 2
  to                   = 20
  user_tag             = "test"
}