provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_fw_server" "test_thunder_fw_server" {
  name         = "test"
  fqdn_name    = "test.com"
  action       = "enable"
  health_check = "ping"
  user_tag     = "testing"
  sampling_enable {
    counters1 = "all"
  }
  port_list {
    port_number  = 30
    protocol     = "tcp"
    action       = "enable"
    health_check = "ping"
    user_tag     = "test_user"
    sampling_enable {
      counters1 = "all"
    }
  }
}