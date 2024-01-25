provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vrrp_a_partition_vrid_all_status_oper" "thunder_vrrp_a_partition_vrid_all_status_oper" {

}
output "get_vrrp_a_partition_vrid_all_status_oper" {
  value = ["${data.thunder_vrrp_a_partition_vrid_all_status_oper.thunder_vrrp_a_partition_vrid_all_status_oper}"]
}