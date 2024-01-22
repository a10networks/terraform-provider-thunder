provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_fw_server_port_oper" "thunder_fw_server_port_oper" {

  name        = "test"
  port_number = 53
  protocol    = "tcp"
}
output "get_fw_server_port_oper" {
  value = ["${data.thunder_fw_server_port_oper.thunder_fw_server_port_oper}"]
}