provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ip_rib_oper" "thunder_ip_rib_oper" {
}
output "get_ip_rib_oper" {
  value = ["${data.thunder_ip_rib_oper.thunder_ip_rib_oper}"]
}