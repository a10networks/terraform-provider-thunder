provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_access_list_extended" "AccessListTest" {
  extd = 100
  rules {
    extd_action              = "permit"
    tcp                      = 1
    src_any                  = 1
    dst_any                  = 1
    ethernet                 = 1
    acl_log                  = 1
    transparent_session_only = 1
  }
  rules {
    extd_action              = "permit"
    tcp                      = 1
    src_any                  = 1
    src_lt                   = 2222
    dst_subnet               = "1.1.1.1"
    dst_mask                 = "255.255.255.0"
    established              = 1
    acl_log                  = 1
    transparent_session_only = 1
  }
}

