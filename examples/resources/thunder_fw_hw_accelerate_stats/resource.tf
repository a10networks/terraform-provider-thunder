provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_fw_hw_accelerate_stats" "thunder_fw_hw_accelerate_stats" {

}
output "get_fw_hw_accelerate_stats" {
  value = ["${data.thunder_fw_hw_accelerate_stats.thunder_fw_hw_accelerate_stats}"]
}