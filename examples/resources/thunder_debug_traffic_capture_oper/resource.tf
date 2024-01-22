provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_debug_traffic_capture_oper" "thunder_debug_traffic_capture_oper" {

}
output "get_debug_traffic_capture_oper" {
  value = ["${data.thunder_debug_traffic_capture_oper.thunder_debug_traffic_capture_oper}"]
}