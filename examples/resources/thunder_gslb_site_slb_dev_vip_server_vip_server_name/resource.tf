provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_site_slb_dev_vip_server_vip_server_name" "thunder_gslb_site_slb_dev_vip_server_vip_server_name" {

  site_name   = "3"
  device_name = "25"
  sampling_enable {
    counters1 = "all"
  }
  vip_name = "vs2"
}