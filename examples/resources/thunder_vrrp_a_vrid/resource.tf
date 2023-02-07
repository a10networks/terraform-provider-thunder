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

resource "thunder_vrrp_a_vrid" "test" {
    vrid_val = 1
    floating_ip {
        ipv6_address_cfg {
            ipv6_address = "2001:1::1" 
        }
        ip_address_cfg {
            ip_address = "10.21.1.1" 
        }
        #ip_address_part_cfg and ipv6_address_part_cfg are private-patition-only
    }
}

resource "thunder_vrrp_a_vrid" "test_l3v_a" {
    provider = thunder.L3V_A
    vrid_val = 2
    floating_ip {
        ipv6_address_part_cfg {
            ipv6_address_partition = "2002:1::1" 
        }
        ip_address_part_cfg {
            ip_address_partition = "10.22.1.1" 
        }
        #ip_address_cfg and ipv6_address_cfg are shared-patition-only
    }
    follow {  
         vrid_lead = "default-vrid-lead" 
    }
}
