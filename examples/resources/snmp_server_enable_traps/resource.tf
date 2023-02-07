provider "thunder" {
    address  = var.dut9049
    username = var.username
    password = var.password
}

provider "thunder" {
    alias = "L3V_A"
    address  = var.dut9049
    username = var.username
    password = var.password
    partition = var.l3v_1
}

resource "thunder_snmp_server_enable_traps" "test00" {
    gslb {
        all = 1
    }
    slb {
      server_up = 1
      server_down = 1
    }
}
