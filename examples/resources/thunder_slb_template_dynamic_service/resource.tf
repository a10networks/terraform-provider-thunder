provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_dynamic_service" "test_thunder_slb_template_dynamic_service" {
  name = "test_service"
  dns_server {
    ipv6_dns_server = "2001:db8:0:200::1"
    ipv4_dns_server = "1.1.1.1"
  }
  user_tag = "test"
}