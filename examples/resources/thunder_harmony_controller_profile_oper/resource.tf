provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_harmony_controller_profile_oper" "thunder_harmony_controller_profile_oper" {


}
output "get_harmony_controller_profile_oper" {
  value = ["${data.thunder_harmony_controller_profile_oper.thunder_harmony_controller_profile_oper}"]
}