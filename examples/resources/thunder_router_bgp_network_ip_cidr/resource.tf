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

resource "thunder_router_bgp" "bgp2000" {
    as_number = 2000
}

resource "thunder_router_bgp_network_ip_cidr" "t01" {
    as_number = thunder_router_bgp.bgp2000.as_number
    network_ipv4_cidr = "19.0.2.0/24"
}

resource "thunder_router_bgp" "bgp3000" {
    provider = thunder.L3V_A
    as_number = 3000
}

resource "thunder_router_bgp_network_ip_cidr" "t02" {
    provider = thunder.L3V_A
    as_number = thunder_router_bgp.bgp3000.as_number
    network_ipv4_cidr = "19.0.3.0/24"
}

