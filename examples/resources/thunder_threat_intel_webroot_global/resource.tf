provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_threat_intel_webroot_global" "thunder_threat_intel_webroot_global" {
  sampling_enable {
    counters1 = "all"
  }
}