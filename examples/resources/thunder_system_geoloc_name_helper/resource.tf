provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_geoloc_name_helper" "thunder_system_geoloc_name_helper" {
  sampling_enable {
    counters1 = "all"
  }
}