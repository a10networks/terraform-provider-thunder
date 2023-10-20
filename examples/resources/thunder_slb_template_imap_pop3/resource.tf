provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_imap_pop3" "test_thunder_slb_template_imap_pop3" {
  name          = "testIMAP"
  starttls      = "enforced"
  logindisabled = 1
  user_tag      = "test"
}