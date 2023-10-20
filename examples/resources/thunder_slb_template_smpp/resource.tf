provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_smpp" "test_thunder_slb_template_smpp" {
  name                    = "testSMPP"
  client_enquire_link     = 1
  server_enquire_link     = 1
  server_enquire_link_val = 60
  user                    = "testUser"
  password                = "testUser"
  user_tag                = "test"
}