provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_persist_stats" "thunder_slb_persist_stats" {
}
output "get_slb_persist_stats" {
  value = ["${data.thunder_slb_persist_stats.thunder_slb_persist_stats}"]
}