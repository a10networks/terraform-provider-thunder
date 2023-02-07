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

resource "thunder_file_bw_list" "BWLIST_1" {
    name = "BWLIST_1"
    protocol = "http"
    host = "192.168.92.200"
    path = "/class-list-files/BWLIST_1.txt"
    use_mgmt_port = 1
}

resource "thunder_file_bw_list" "L3V_BWLIST_1" {
    provider = thunder.L3V_A
    name = "BWLIST_1"
    protocol = "https"
    host = "192.168.92.200"
    path = "/class-list-files/BWLIST_1.txt"
    use_mgmt_port = 1
}

