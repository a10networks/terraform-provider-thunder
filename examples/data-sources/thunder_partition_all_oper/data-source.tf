provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_partition_all_oper" "thunder_partition_all_oper" {
}
output "get_partition_all_oper" {
  value = ["${data.thunder_partition_all_oper.thunder_partition_all_oper}"]
}