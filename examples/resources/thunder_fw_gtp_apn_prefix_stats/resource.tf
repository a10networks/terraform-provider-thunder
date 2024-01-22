provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_fw_gtp_apn_prefix_stats" "thunder_fw_gtp_apn_prefix_stats" {
}
output "get_fw_gtp_apn_prefix_stats" {
  value = ["${data.thunder_fw_gtp_apn_prefix_stats.thunder_fw_gtp_apn_prefix_stats}"]
}