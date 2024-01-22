provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_site_ip_server" "thunder_gslb_site_ip_server" {

  site_name      = "3"
  ip_server_name = "test-server1"

}