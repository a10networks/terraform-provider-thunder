provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_detection_statistics_stats" "thunder_ddos_detection_statistics_stats" {
}
output "get_ddos_detection_statistics_stats" {
  value = ["${data.thunder_ddos_detection_statistics_stats.thunder_ddos_detection_statistics_stats}"]
}