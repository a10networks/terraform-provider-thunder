provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_persist_source_ip" "persist-source" {
  name                    = "persistip1"
  incl_dst_ip             = 1
  timeout                 = 100
  match_type              = 1
  dont_honor_conn_rules   = 1
  enforce_higher_priority = 1
  hash_persist            = 1
  incl_sport              = 1
  netmask                 = "/23"
  netmask6                = 96
  server                  = 1
  primary_port            = 443
  scan_all_members        = 1
}