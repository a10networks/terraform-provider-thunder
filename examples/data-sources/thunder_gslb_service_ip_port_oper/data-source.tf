provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_service_ip_port_oper" "thunder_gslb_service_ip_port_oper" {

  node_name  = "vs2"
  port_num   = 30672
  port_proto = "tcp"
}
output "get_gslb_service_ip_port_oper" {
  value = ["${data.thunder_gslb_service_ip_port_oper.thunder_gslb_service_ip_port_oper}"]
}