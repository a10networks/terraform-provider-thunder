provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

data "thunder_system_cpu_ctrl_cpu_oper" "test"{
}

output "testing" {
  value = ["${data.thunder_system_cpu_ctrl_cpu_oper.test}"]
}