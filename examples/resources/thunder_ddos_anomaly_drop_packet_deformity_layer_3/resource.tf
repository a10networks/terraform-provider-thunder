provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_anomaly_drop_packet_deformity_layer_3" "thunder_ddos_anomaly_drop_packet_deformity_layer_3" {
  capture_config = "10"
  log            = 1
  toggle         = "disable"
}