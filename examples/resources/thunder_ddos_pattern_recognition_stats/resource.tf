provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_pattern_recognition_stats" "thunder_ddos_pattern_recognition_stats" {
}
output "get_ddos_pattern_recognition_stats" {
  value = ["${data.thunder_ddos_pattern_recognition_stats.thunder_ddos_pattern_recognition_stats}"]
}