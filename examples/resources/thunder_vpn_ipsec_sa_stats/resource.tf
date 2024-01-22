provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vpn_ipsec_sa_stats" "thunder_vpn_ipsec_sa_stats" {
  sampling_enable {
    counters1 = "all"
  }

}
output "get_vpn_ipsec_sa_stats" {
  value = ["${data.thunder_vpn_ipsec_sa_stats.thunder_vpn_ipsec_sa_stats}"]
}