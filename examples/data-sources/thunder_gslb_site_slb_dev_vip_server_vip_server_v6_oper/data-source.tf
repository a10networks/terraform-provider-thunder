provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_site_slb_dev_vip_server_vip_server_v6_oper" "thunder_gslb_site_slb_dev_vip_server_vip_server_v6_oper" {

  site_name   = "3"
  device_name = "25"
  ipv6        = "cf58:9ec7:91f2:cd5c:219c:66eb:d05b:51ca"

}
output "get_gslb_site_slb_dev_vip_server_vip_server_v6_oper" {
  value = ["${data.thunder_gslb_site_slb_dev_vip_server_vip_server_v6_oper.thunder_gslb_site_slb_dev_vip_server_vip_server_v6_oper}"]
}