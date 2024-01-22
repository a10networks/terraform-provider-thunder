provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_scaleout_debug_cluster_database_oper" "thunder_scaleout_debug_cluster_database_oper" {
}
output "get_scaleout_debug_cluster_database_oper" {
  value = ["${data.thunder_scaleout_debug_cluster_database_oper.thunder_scaleout_debug_cluster_database_oper}"]
}