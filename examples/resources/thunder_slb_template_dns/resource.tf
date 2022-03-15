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

resource "thunder_class_list" "clist_a" {
    name = "CLIST_A"
    type = "dns"
}

resource "thunder_slb_template_dns" "t_dns_01" {
    name = "t_dns_01"
    class_list {
        name = thunder_class_list.clist_a.name
        lid_list {
            lidnum = 10
            conn_rate_limit = 3000
            per = 10
            over_limit_action = 1
            action_value = "dns-cache-enable"
            lockout = 30
            log = 0
            dns {
                cache_action = "cache-disable"
                honor_server_response_ttl = 1
            }
        }
    }
}

resource "thunder_class_list" "clist_b" {
    provider = thunder.L3V_A
    name = "CLIST_B"
    type = "dns"
}

resource "thunder_slb_template_dns" "t_dns_02" {
    provider = thunder.L3V_A
    name = "t_dns_02"
    class_list {
        name = thunder_class_list.clist_b.name
        lid_list {
            lidnum = 10
            conn_rate_limit = 3000
            per = 10
            over_limit_action = 1
            action_value = "dns-cache-enable"
            lockout = 30
            log = 0
            dns {
                cache_action = "cache-disable"
                honor_server_response_ttl = 1
            }
        }
    }
}
