provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_ddos_protection" "thunder_fw_ddos_protection" {
  action {
    action_type = "drop"
    expiration  = 5
  }
  dynamic_blacklist {
    dynamic_blacklist_action = "enable"
    dir                      = "both"
  }
  sampling_enable {
    counters1 = "all"
  }
}