provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_link_probe_entry_oper" "thunder_slb_link_probe_entry_oper" {
}
output "get_slb_link_probe_entry_oper" {
  value = ["${data.thunder_slb_link_probe_entry_oper.thunder_slb_link_probe_entry_oper}"]
}