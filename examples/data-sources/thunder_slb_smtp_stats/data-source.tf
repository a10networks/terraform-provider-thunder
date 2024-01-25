provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_smtp_stats" "thunder_slb_smtp_stats" {
}
output "get_slb_smtp_stats" {
  value = ["${data.thunder_slb_smtp_stats.thunder_slb_smtp_stats}"]
}