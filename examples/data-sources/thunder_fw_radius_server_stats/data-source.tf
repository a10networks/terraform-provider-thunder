provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_fw_radius_server_stats" "thunder_fw_radius_server_stats" {

}
output "get_fw_radius_server_stats" {
  value = ["${data.thunder_fw_radius_server_stats.thunder_fw_radius_server_stats}"]
}