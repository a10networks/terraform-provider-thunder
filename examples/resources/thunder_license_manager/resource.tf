provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_license_manager" "thunder_license_manager" {
  bandwidth_base         = 23443
  bandwidth_unrestricted = 1
  connect {
    connect = 0
  }
  host_list {
    host_ipv4 = "10.0.1.2"
    host_ipv6 = "7b5c:2272:a2a7:94c7:4004:8cbc:666f:06f1"
    port      = 443
  }
  instance_name = "1.4.7.6"
  interval      = 1
  overage {
    days    = 277
    hours   = 3
    minutes = 28
    seconds = 32
    gb      = 49686
    mb      = 219
    kb      = 646
    bytes   = 1019
  }
  reminder_list {
    reminder_value = 5885
  }
  sn = "22"
}