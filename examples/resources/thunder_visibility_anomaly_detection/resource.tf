provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_anomaly_detection" "thunder_visibility_anomaly_detection" {
  sensitivity                 = "low"
  restart_learning_on_anomaly = 1
  feature_status              = "disable"
  logging                     = "disable"
}