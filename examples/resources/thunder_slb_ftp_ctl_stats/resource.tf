provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_ftp_ctl_stats" "thunder_slb_ftp_ctl_stats" {
}
output "get_slb_ftp_ctl_stats" {
  value = ["${data.thunder_slb_ftp_ctl_stats.thunder_slb_ftp_ctl_stats}"]
}