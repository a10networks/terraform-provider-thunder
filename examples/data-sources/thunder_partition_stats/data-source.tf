provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_partition_stats" "thunder_partition_stats" {

  partition_name = "test"
}
output "get_partition_stats" {
  value = ["${data.thunder_partition_stats.thunder_partition_stats}"]
}