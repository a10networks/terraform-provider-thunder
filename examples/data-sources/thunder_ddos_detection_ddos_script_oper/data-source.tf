provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_detection_ddos_script_oper" "thunder_ddos_detection_ddos_script_oper" {
}
output "get_ddos_detection_ddos_script_oper" {
  value = ["${data.thunder_ddos_detection_ddos_script_oper.thunder_ddos_detection_ddos_script_oper}"]
}