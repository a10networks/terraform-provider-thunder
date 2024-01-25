provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_lsn_radius_server_oper" "thunder_cgnv6_lsn_radius_server_oper" {
  oper {
    radius_table_entries_list {
      inside_ip   = "10.10.10.10"
      inside_ipv6 = "2001:db8:3333:4444:5555:6666:7777:8888"
    }
  }
}
output "get_cgnv6_lsn_radius_server_oper" {
  value = ["${data.thunder_cgnv6_lsn_radius_server_oper.thunder_cgnv6_lsn_radius_server_oper}"]
}