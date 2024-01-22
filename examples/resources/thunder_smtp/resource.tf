provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_smtp" "thunder_smtp" {
  mailfrom           = "45"
  needauthentication = 0
  port               = 45061
  smtp_server        = "160"
  username_cfg {
    username = "3"
    password {
      smtp_password = "20"
    }
  }
}