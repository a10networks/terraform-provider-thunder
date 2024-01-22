provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_site_ip_server_oper" "thunder_gslb_site_ip_server_oper" {

  site_name      = "3"
  ip_server_name = "test-server1"

}
output "get_gslb_site_ip_server_oper" {
  value = ["${data.thunder_gslb_site_ip_server_oper.thunder_gslb_site_ip_server_oper}"]
}