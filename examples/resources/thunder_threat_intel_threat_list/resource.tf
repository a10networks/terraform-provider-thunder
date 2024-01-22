provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_threat_intel_threat_list" "thunder_threat_intel_threat_list" {
  name           = "test_threat"
  type           = "webroot"
  all_categories = 1

  sampling_enable {
    counters1 = "all"
  }
  user_tag = "2"
}