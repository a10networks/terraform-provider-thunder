provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_bootimage_oper" "thunder_bootimage_oper" {
}
output "get_bootimage_oper" {
  value = ["${data.thunder_bootimage_oper.thunder_bootimage_oper}"]
}