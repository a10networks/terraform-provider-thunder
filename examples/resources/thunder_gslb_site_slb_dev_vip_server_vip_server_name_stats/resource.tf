provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_site_slb_dev_vip_server_vip_server_name_stats" "thunder_gslb_site_slb_dev_vip_server_vip_server_name_stats" {

  site_name   = "3"
  device_name = "25"
  vip_name    = "vs2"
}
output "get_gslb_site_slb_dev_vip_server_vip_server_name_stats" {
  value = ["${data.thunder_gslb_site_slb_dev_vip_server_vip_server_name_stats.thunder_gslb_site_slb_dev_vip_server_vip_server_name_stats}"]
}