provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

data "thunder_clock_show_oper" "test"{
}

output "testing" {
  value = ["${data.thunder_clock_show_oper.test}"]
}