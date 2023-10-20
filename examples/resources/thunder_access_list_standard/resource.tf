provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_access_list_standard" "AccessListTest" {
  std = 50
  stdrules {
    seq_num                  = 1
    action                   = "permit"
    host                     = "192.168.40.1"
    log                      = 1
    transparent_session_only = 1
  }
  stdrules {
    seq_num                  = 2
    action                   = "permit"
    subnet                   = "192.168.0.0"
    rev_subnet_mask          = "255.255.0.0"
    log                      = 1
    transparent_session_only = 1
  }
}

