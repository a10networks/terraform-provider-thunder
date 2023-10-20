provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

data "thunder_system_throughput_stats" "throughput" {}

output "get_throughput" {
  value = ["${data.thunder_system_throughput_stats.throughput}"]
}