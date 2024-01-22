provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_monitor_method_rtsp" "thunder_health_monitor_method_rtsp" {

  name      = "tf_test"
  rtsp      = 1
  rtsp_port = 55412
  rtspurl   = "/www.example.com"
}