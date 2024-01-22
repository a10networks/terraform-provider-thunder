provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_fw_alg_esp_stats" "thunder_fw_alg_esp_stats" {

}
output "get_fw_alg_esp_stats" {
  value = ["${data.thunder_fw_alg_esp_stats.thunder_fw_alg_esp_stats}"]
}