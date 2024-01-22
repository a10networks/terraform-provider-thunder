provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_geoloc" "thunder_system_geoloc" {
  sampling_enable {
    counters1 = "all"
  }
}