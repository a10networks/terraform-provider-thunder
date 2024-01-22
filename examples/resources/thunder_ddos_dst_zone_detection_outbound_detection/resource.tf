provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_detection_outbound_detection" "thunder_ddos_dst_zone_detection_outbound_detection" {

  configuration = "configuration"
  indicator_list {
    type          = "pkt-rate"
    threshold_num = 858137122
    user_tag      = "33"
  }
  toggle = "disable"
}