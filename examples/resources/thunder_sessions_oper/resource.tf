provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

data "thunder_sessions_oper" "test"{ }

output "get_sessions" {
  value = ["${data.thunder_sessions_oper.test}"]
}