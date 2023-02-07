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

resource "thunder_interface_ve" "test00" {
    ifnum = 300
    ip {
        address_list {
            ipv4_address = "56.56.56.56"
            ipv4_netmask = "/16"
        }
    }
}

resource "thunder_interface_ve" "test01" {
    provider = thunder.L3V_A
    ifnum = 400
    ip {
        address_list {
            ipv4_address = "56.156.0.1"
            ipv4_netmask = "/24"
        }
    }
}

resource "thunder_interface_ve" "test02" {
    provider = thunder.L3V_A
    ifnum = 401
    ipv6 {
        address_list {
            ipv6_addr = "fd08::222/64"
        }
    }
}
