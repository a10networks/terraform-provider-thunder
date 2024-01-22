provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_fw_tcp_syn_cookie_stats" "thunder_fw_tcp_syn_cookie_stats" {

}
output "get_fw_tcp_syn_cookie_stats" {
  value = ["${data.thunder_fw_tcp_syn_cookie_stats.thunder_fw_tcp_syn_cookie_stats}"]
}