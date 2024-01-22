provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_anomaly_drop_packet_deformity_layer_4" "thunder_ddos_anomaly_drop_packet_deformity_layer_4" {

  capture_config = "10"
  log            = 1
  toggle         = "disable"

}