provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_site_slb_dev_vip_server_vip_server_v4_oper" "thunder_gslb_site_slb_dev_vip_server_vip_server_v4_oper" {

  site_name   = "3"
  device_name = "25"
  ipv4        = "10.10.10.10"
}
output "get_gslb_site_slb_dev_vip_server_vip_server_v4_oper" {
  value = ["${data.thunder_gslb_site_slb_dev_vip_server_vip_server_v4_oper.thunder_gslb_site_slb_dev_vip_server_vip_server_v4_oper}"]
}