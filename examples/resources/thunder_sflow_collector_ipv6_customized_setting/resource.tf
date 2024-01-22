provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sflow_collector_ipv6_customized_setting" "thunder_sflow_collector_ipv6_customized_setting" {

  a10_proprietary_polling = 0
  counter_polling         = 0
  event_notification      = 0
  export_enable           = "export"
  packet_sampling         = 0
  addr                    = "2001:db8:3333:4444:5555:6666:7777:8888"
  port                    = 4428
}