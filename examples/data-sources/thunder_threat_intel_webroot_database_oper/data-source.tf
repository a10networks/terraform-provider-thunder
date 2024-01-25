provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_threat_intel_webroot_database_oper" "thunder_threat_intel_webroot_database_oper" {
}
output "get_threat_intel_webroot_database_oper" {
  value = ["${data.thunder_threat_intel_webroot_database_oper.thunder_threat_intel_webroot_database_oper}"]
}