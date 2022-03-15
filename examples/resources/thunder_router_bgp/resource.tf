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

resource "thunder_router_bgp" "bgp1" {
    as_number = 101
    neighbor {
        ipv4_neighbor_list {   
            neighbor_ipv4 = "10.1.1.104" 
            activate =  1 
            nbr_remote_as = 104
            allowas_in =  1
            allowas_in_count =  10 
            graceful_restart =  1
        }
    }
}

resource "thunder_router_bgp" "bgp2" {
    provider = thunder.L3V_A
    as_number = 201
    neighbor {
        ipv4_neighbor_list {   
            neighbor_ipv4 = "10.1.1.204" 
            activate =  1 
            nbr_remote_as = 204
            allowas_in =  1
            allowas_in_count =  10 
            graceful_restart =  1
        }
    }
}
