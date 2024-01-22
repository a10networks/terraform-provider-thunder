provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vrrp_a_hostid_oper" "thunder_vrrp_a_hostid_oper" {
}
output "get_vrrp_a_hostid_oper" {
  value = ["${data.thunder_vrrp_a_hostid_oper.thunder_vrrp_a_hostid_oper}"]
}