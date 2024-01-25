provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_interface_lif_oper" "thunder_interface_lif_oper" {

  ifname = "test"
}
output "get_interface_lif_oper" {
  value = ["${data.thunder_interface_lif_oper.thunder_interface_lif_oper}"]
}