provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_site_slb_dev_vip_server_vip_server_v4" "thunder_gslb_site_slb_dev_vip_server_vip_server_v4" {

  site_name   = "3"
  device_name = "25"
  ipv4        = "10.10.10.10"
  sampling_enable {
    counters1 = "all"
  }
}