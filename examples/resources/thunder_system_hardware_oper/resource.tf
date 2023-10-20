provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

data "thunder_system_hardware_oper" "test"{}

output "get_disk" {
  value = ["${data.thunder_system_hardware_oper.test}"]
}