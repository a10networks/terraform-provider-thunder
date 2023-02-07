provider "thunder" {
    address  = var.dut90192
    username = var.username
    password = var.password
}

provider "thunder" {
    alias = "L3V_A"
    address  = var.dut90192
    username = var.username
    password = var.password
    partition = var.l3v_1
}

resource "thunder_snmp_server_disable_traps" "S_L3V_A" {
    provider = thunder.L3V_A
    all = 0
    gslb = 0
    slb = 1
    slb_change = 1
    snmp = 1
    vrrp_a = 1
}
