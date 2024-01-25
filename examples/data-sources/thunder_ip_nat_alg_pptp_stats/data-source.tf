provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ip_nat_alg_pptp_stats" "thunder_ip_nat_alg_pptp_stats" {

}
output "get_ip_nat_alg_pptp_stats" {
  value = ["${data.thunder_ip_nat_alg_pptp_stats.thunder_ip_nat_alg_pptp_stats}"]
}