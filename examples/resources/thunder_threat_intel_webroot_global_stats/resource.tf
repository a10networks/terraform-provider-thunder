provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_threat_intel_webroot_global_stats" "thunder_threat_intel_webroot_global_stats" {
}
output "get_threat_intel_webroot_global_stats" {
  value = ["${data.thunder_threat_intel_webroot_global_stats.thunder_threat_intel_webroot_global_stats}"]
}