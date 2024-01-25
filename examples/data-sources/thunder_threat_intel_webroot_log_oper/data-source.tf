provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_threat_intel_webroot_log_oper" "thunder_threat_intel_webroot_log_oper" {

}
output "get_threat_intel_webroot_log_oper" {
  value = ["${data.thunder_threat_intel_webroot_log_oper.thunder_threat_intel_webroot_log_oper}"]
}