provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_service_port_oper" "thunder_gslb_service_port_oper" {
  oper {
    service_port_list {
      active_real_server  = active_real_server
      current_connections = current_connections
    }
  }
}
output "get_gslb_service_port_oper" {
  value = ["${data.thunder_gslb_service_port_oper.thunder_gslb_service_port_oper}"]
}