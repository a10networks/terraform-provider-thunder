provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_service_ip_port_stats" "thunder_gslb_service_ip_port_stats" {

  node_name  = "vs2"
  port_num   = 30672
  port_proto = "tcp"
}
output "get_gslb_service_ip_port_stats" {
  value = ["${data.thunder_gslb_service_ip_port_stats.thunder_gslb_service_ip_port_stats}"]
}