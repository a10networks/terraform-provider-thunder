provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_server_port_stats" "thunder_slb_server_port_stats" {

  name        = "test-server1"
  port_number = 30715
  protocol    = "tcp"

}
output "get_slb_server_port_stats" {
  value = ["${data.thunder_slb_server_port_stats.thunder_slb_server_port_stats}"]
}