provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_src_entry_l4_type" "thunder_ddos_src_entry_l4_type" {
  src_entry_name = "25"
  action         = "permit"
  protocol       = "tcp"
  user_tag       = "100"
}