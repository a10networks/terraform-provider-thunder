provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_fw_server_stats" "thunder_fw_server_stats" {

  name = "test"
}
output "get_fw_server_stats" {
  value = ["${data.thunder_fw_server_stats.thunder_fw_server_stats}"]
}