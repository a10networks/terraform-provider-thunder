provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dst_entry_stats" "thunder_ddos_dst_entry_stats" {
  dst_entry_name = "test"

}
output "get_ddos_dst_entry_stats" {
  value = ["${data.thunder_ddos_dst_entry_stats.thunder_ddos_dst_entry_stats}"]
}