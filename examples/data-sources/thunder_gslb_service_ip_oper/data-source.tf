provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_service_ip_oper" "thunder_gslb_service_ip_oper" {

  node_name = "s27"
}
output "get_gslb_service_ip_oper" {
  value = ["${data.thunder_gslb_service_ip_oper.thunder_gslb_service_ip_oper}"]
}