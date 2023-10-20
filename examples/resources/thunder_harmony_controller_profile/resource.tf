provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_harmony_controller_profile" "profile" {
  host              = "192.0.2.65"
  port              = 8443
  user_name         = "terraform@a10networks.com"
  secret_value      = "admin123"
  provider2         = "root"
  action            = "register"
  use_mgmt_port     = 1
  region            = "US/WEST"
  availability_zone = "DC-1"
  thunder_mgmt_ip {
    ip_address = "192.0.2.65"
  }
}