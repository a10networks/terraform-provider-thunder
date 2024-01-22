provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sflow_collector_ip_customized_setting" "thunder_sflow_collector_ip_customized_setting" {
  a10_proprietary_polling = 0
  counter_polling         = 0
  event_notification      = 0
  export_enable           = "export"
  packet_sampling         = 0
  addr                    = "10.10.10.10"
  port                    = 38248
}