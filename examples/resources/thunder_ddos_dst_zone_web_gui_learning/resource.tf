provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_web_gui_learning" "thunder_ddos_dst_zone_web_gui_learning" {
  zone_name     = "testZone"
  duration      = "6hour"
  starting_time = "2"

}