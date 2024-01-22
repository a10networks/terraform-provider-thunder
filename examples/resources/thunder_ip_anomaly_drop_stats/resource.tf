provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ip_anomaly_drop_stats" "thunder_ip_anomaly_drop_stats" {
}
output "get_ip_anomaly_drop_stats" {
  value = ["${data.thunder_ip_anomaly_drop_stats.thunder_ip_anomaly_drop_stats}"]
}