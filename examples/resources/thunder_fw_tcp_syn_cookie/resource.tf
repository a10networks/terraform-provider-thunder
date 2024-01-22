provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_tcp_syn_cookie" "thunder_fw_tcp_syn_cookie" {
  sampling_enable {
    counters1 = "all"
  }
  syn_cookie_enable       = 1
  syn_cookie_on_threshold = 599320
  syn_cookie_on_timeout   = 130
}