provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_link_probe_entry_stats" "thunder_slb_link_probe_entry_stats" {
}
output "get_slb_link_probe_entry_stats" {
  value = ["${data.thunder_slb_link_probe_entry_stats.thunder_slb_link_probe_entry_stats}"]
}