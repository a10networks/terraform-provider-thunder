provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_ssl_expire_check" "Slb_Ssl_Expire_Check_Exception" {
  exception {
    action           = "add"
    certificate_name = "SLBCERT"
  }
}
resource "thunder_slb_ssl_expire_check" "Slb_Ssl_Expire_Check_Email" {
  ssl_expire_email_address = "ssltest@a10networks.com"
  interval_days            = 2
  expire_address1          = "testssl@incedoinc.com"
  before                   = 30
}