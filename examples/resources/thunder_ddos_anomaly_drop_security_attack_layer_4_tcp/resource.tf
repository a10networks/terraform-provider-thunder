provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_anomaly_drop_security_attack_layer_4_tcp" "thunder_ddos_anomaly_drop_security_attack_layer_4_tcp" {
  capture_config = "10"
  log            = 1
  toggle         = "disable"
}