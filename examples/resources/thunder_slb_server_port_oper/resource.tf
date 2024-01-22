provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_server_port_oper" "thunder_slb_server_port_oper" {

  port_number = 30715
  protocol    = "tcp"
  name        = "test-server1"

}
output "get_slb_server_port_oper" {
  value = ["${data.thunder_slb_server_port_oper.thunder_slb_server_port_oper}"]
}