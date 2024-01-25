provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_plat_cpu_drop_oper" "thunder_plat_cpu_drop_oper" {
}
output "get_plat_cpu_drop_oper" {
  value = ["${data.thunder_plat_cpu_drop_oper.thunder_plat_cpu_drop_oper}"]
}