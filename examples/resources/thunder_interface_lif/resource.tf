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

resource "thunder_interface_lif" "test00" {
    ifname = 2
    ip {
        address_list {
            ipv4_address = "56.56.56.56"
            ipv4_netmask = "/16"
        }
    }
}

resource "thunder_interface_lif" "test01" {
    provider = thunder.L3V_A
    ifname = 7
    ip {
        address_list {
            ipv4_address = "56.156.0.1"
            ipv4_netmask = "/24"
        }
    }
}

resource "thunder_interface_lif" "test02" {
    provider = thunder.L3V_A
    ifname = 8
    ipv6 {
        address_list {
            ipv6_addr = "fd08::222/64"
        }
    }
}
