provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_event_notification_kafka_server" "thunder_event_notification_kafka_server" {
  port = 9092
  sampling_enable {
    counters1 = "all"
  }
  use_mgmt_port = 0
}