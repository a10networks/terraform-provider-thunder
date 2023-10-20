provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

data "thunder_slb_server_oper" "test"{
  name = "server1"
}

output "testing" {
  value = ["${data.thunder_slb_server_oper.test}"]
}