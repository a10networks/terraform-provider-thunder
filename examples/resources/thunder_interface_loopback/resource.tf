provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

provider "thunder" {
  alias     = "L3V_A"
  address   = var.dut9049
  username  = var.username
  password  = var.password
  partition = var.l3v_1
}

resource "thunder_interface_loopback" "test00" {
  ifnum = 0
  ip {
    address_list {
      ipv4_address = "10.127.0.11"
      ipv4_netmask = "/24"
    }
  }
}

resource "thunder_interface_loopback" "test02" {
  provider = thunder.L3V_A
  ifnum    = 2
  ip {
    address_list {
      ipv4_address = "10.127.2.11"
      ipv4_netmask = "/17"
    }
  }
}

