provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_partition_available_id_oper" "thunder_partition_available_id_oper" {
}
output "get_partition_available_id_oper" {
  value = ["${data.thunder_partition_available_id_oper.thunder_partition_available_id_oper}"]
}