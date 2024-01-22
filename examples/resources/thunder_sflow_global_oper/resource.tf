provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_sflow_global_oper" "thunder_sflow_global_oper" {
}
output "get_sflow_global_oper" {
  value = ["${data.thunder_sflow_global_oper.thunder_sflow_global_oper}"]
}