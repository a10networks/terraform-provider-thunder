provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ntp_status_oper" "thunder_ntp_status_oper" {

}
output "get_ntp_status_oper" {
  value = ["${data.thunder_ntp_status_oper.thunder_ntp_status_oper}"]
}