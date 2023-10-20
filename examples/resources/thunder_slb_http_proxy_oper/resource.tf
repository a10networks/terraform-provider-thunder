provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

data "thunder_slb_http_proxy_oper" "test"{
}

output "testing" {
  value = ["${data.thunder_slb_http_proxy_oper.test}"]
}