provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system" "test" {
  promiscuous_mode = 1
  anomaly_log      = 1
  attack_log       = 0
  session {
    sampling_enable {
      counters1 = "all"
    }
  }
}