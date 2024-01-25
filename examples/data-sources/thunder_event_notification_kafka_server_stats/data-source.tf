provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_event_notification_kafka_server_stats" "thunder_event_notification_kafka_server_stats" {
}
output "get_event_notification_kafka_server_stats" {
  value = ["${data.thunder_event_notification_kafka_server_stats.thunder_event_notification_kafka_server_stats}"]
}