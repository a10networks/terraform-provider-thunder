provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_monitor_sflow" "thunder_visibility_monitor_sflow" {
  listening_port = 63

}