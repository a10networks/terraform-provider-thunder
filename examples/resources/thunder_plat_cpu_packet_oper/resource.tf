provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

data "thunder_plat_cpu_packet_oper" "test"{
}

output "get_plat_cpu_packet" {
  value = ["${data.thunder_plat_cpu_packet_oper.test}"]
}