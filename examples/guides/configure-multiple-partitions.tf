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

resource "thunder_virtual_server" "VS01" {
    name = "VS1+1"
    ip_address = "10.22.0.1"
}

resource "thunder_virtual_server" "test02" {
    name = "VS1:2"
    ip_address = "10.22.0.2"
}

resource "thunder_virtual_server" "test03" {
    provider = thunder.L3V_A
    name = "VS1-3"
    ip_address = "10.22.19.3"
}
