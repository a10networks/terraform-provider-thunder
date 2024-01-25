provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_sctp_global_stats" "thunder_sctp_global_stats" {
}
output "get_sctp_global_stats" {
  value = ["${data.thunder_sctp_global_stats.thunder_sctp_global_stats}"]
}