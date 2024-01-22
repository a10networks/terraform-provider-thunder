provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_site_slb_dev_vip_server_vip_server_v4_stats" "thunder_gslb_site_slb_dev_vip_server_vip_server_v4_stats" {
  ipv4 = "10.10.10.10"
  stats {
    dev_vip_hits   = dev_vip_hits
    dev_vip_recent = dev_vip_recent
  }
}
output "get_gslb_site_slb_dev_vip_server_vip_server_v4_stats" {
  value = ["${data.thunder_gslb_site_slb_dev_vip_server_vip_server_v4_stats.thunder_gslb_site_slb_dev_vip_server_vip_server_v4_stats}"]
}