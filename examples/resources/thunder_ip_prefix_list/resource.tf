provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_ip_prefix_list" "IpPrefixList1" {
  name = "PREFIX-LIST"
  rules {
    le          = 32
    ge          = 24
    description = "Testing_resource"
    any         = 1
    action      = "deny"
    ipaddr      = "2.2.4.7/8"
    seq         = 1
  }
}