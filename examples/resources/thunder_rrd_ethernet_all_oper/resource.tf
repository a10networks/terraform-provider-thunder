provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_rrd_ethernet_all_oper" "thunder_rrd_ethernet_all_oper" {
}
output "get_rrd_ethernet_all_oper" {
  value = ["${data.thunder_rrd_ethernet_all_oper.thunder_rrd_ethernet_all_oper}"]
}