provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_ipv6_ospf" "thunder_debug_ipv6_ospf" {
  all {
    dumy = 0
  }
  bfd {
    dumy = 0
  }
  dumy = 0
  events {
    abr    = 0
    asbr   = 0
    os     = 0
    router = 0
    vlink  = 0
  }
  ifsm {
    events = 0
    status = 0
    timers = 0
  }
  lsa {
    flooding = 0
    gererate = 0
    install  = 0
    maxage   = 0
    refresh  = 0
  }
  nfsm {
    events = 0
    status = 0
    timers = 0
  }
  nsm {
    interface    = 0
    redistribute = 0
  }
  packet {
    dd         = 0
    detail     = 0
    hello      = 0
    ls_ack     = 0
    ls_request = 0
    ls_update  = 0
    recv       = 0
    send       = 0
  }
  route {
    ase     = 0
    ia      = 0
    install = 0
    spf     = 0
  }
}