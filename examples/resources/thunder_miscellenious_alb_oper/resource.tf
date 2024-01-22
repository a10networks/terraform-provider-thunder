provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_miscellenious_alb_oper" "thunder_miscellenious_alb_oper" {
}
output "get_miscellenious_alb_oper" {
  value = ["${data.thunder_miscellenious_alb_oper.thunder_miscellenious_alb_oper}"]
}