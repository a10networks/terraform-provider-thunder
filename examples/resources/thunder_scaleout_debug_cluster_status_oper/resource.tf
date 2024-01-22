provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_scaleout_debug_cluster_status_oper" "thunder_scaleout_debug_cluster_status_oper" {
}
output "get_scaleout_debug_cluster_status_oper" {
  value = ["${data.thunder_scaleout_debug_cluster_status_oper.thunder_scaleout_debug_cluster_status_oper}"]
}