provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_fw_full_cone_session_oper" "thunder_fw_full_cone_session_oper" {

}
output "get_fw_full_cone_session_oper" {
  value = ["${data.thunder_fw_full_cone_session_oper.thunder_fw_full_cone_session_oper}"]
}