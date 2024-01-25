provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_event_notification_kafka_server_oper" "thunder_event_notification_kafka_server_oper" {
}
output "get_event_notification_kafka_server_oper" {
  value = ["${data.thunder_event_notification_kafka_server_oper.thunder_event_notification_kafka_server_oper}"]
}