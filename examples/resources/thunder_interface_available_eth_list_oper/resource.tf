provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

data "thunder_interface_available_eth_list_oper" "test"{}

output "get_disk" {
  value = ["${data.thunder_interface_available_eth_list_oper.test}"]
}