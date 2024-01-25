provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_rrd_memory_oper" "thunder_rrd_memory_oper" {
}
output "get_rrd_memory_oper" {
  value = ["${data.thunder_rrd_memory_oper.thunder_rrd_memory_oper}"]
}