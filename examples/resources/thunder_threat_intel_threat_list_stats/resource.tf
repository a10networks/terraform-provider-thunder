provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_threat_intel_threat_list_stats" "thunder_threat_intel_threat_list_stats" {

  name = "test_threat"
}
output "get_threat_intel_threat_list_stats" {
  value = ["${data.thunder_threat_intel_threat_list_stats.thunder_threat_intel_threat_list_stats}"]
}